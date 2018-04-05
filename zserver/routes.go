package zserver

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"../logger"
	"../plugins"
	"../plugins/add"
	"../plugins/incrementByStep"
	"../plugins/multiply"
	"../utils"
)

func defaultHandler(c *gin.Context) {
	c.String(http.StatusOK, "Working")
}

func genericResponseErrorHandler(c *gin.Context) {
	r := recover()
	if r != nil {
		if err, ok := r.(*utils.CustomError); ok {
			logger.Errorf("HTTPResponse error (%d) for URL '%s': %s", err.ErrorCode(), c.Request.URL.Path, err.Error())
			c.String(err.ErrorCode(), err.Error())
		} else {
			logger.Errorf("HTTPResponse error (%d) for URL '%s': %s", http.StatusInternalServerError, c.Request.URL.Path, r.(string))
			c.String(http.StatusInternalServerError, r.(string))
		}
	}
}

func jsonHandler(c *gin.Context) {
	response := make(map[string]string)
	response["0"] = "abcd"
	c.JSON(http.StatusOK, response)
}

func getChannelConfig(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Config)
}

func setChannelConfig(c *gin.Context) {
	defer genericResponseErrorHandler(c)

	var pluginsConfig interface{}
	c.BindJSON(&pluginsConfig)
	err := utils.UpdateConfig(pluginsConfig)
	if err != nil {
		panic(utils.NewCustomError(err.Error(), http.StatusInternalServerError))
	}
	c.String(http.StatusOK, "")
}

func getPluginsConfig(c *gin.Context) {
	c.JSON(http.StatusOK, utils.PluginsConfig)
}

func getQueryParam(q map[string][]string, key string) string {
	val, ok := q[key]
	if !ok {
		panic(utils.NewCustomError(fmt.Sprintf("'%s' is a required integer parameter.", key), http.StatusBadRequest))
	}
	return val[0]
}

func executePlugin(pluginName string, args ...interface{}) ([]interface{}, error) {
	defer utils.GenericErrorHandler(nil)

	pluginReturnValue := plugins.PluginReturnValue{}

	switch pluginName {
	case "incrementByStep":
		incrementByStepPlugin, err := incrementByStep.New(args...)
		if err != nil {
			utils.SafeErrorHandler(err)
			return pluginReturnValue.Values, err
		}
		pluginReturnValue, err = incrementByStepPlugin.Execute()
		return pluginReturnValue.Values, err
	case "add":
		addPlugin, err := add.New(args...)
		if err != nil {
			utils.SafeErrorHandler(err)
			return pluginReturnValue.Values, err
		}
		pluginReturnValue, err = addPlugin.Execute()
		return pluginReturnValue.Values, err
	case "multiply":
		multiplyPlugin, err := multiply.New(args...)
		if err != nil {
			utils.SafeErrorHandler(err)
			return pluginReturnValue.Values, err
		}
		pluginReturnValue, err = multiplyPlugin.Execute()
		return pluginReturnValue.Values, err
	default:
		{
			panic(utils.NewCustomError(fmt.Sprintf("Step '%s' is not supported.", pluginName), http.StatusNotImplemented))
		}
	}
}

func upload(c *gin.Context) {
	defer genericResponseErrorHandler(c)

	var err error

	q := c.Request.URL.Query()
	channel := getQueryParam(q, "channel")
	if channel == "" {
		panic(utils.NewCustomError("'channel' is a required string parameter.", http.StatusBadRequest))
	}
	channelConfig, ok := utils.Config.Channels[channel]
	if !ok {
		panic(utils.NewCustomError(fmt.Sprintf("'%s' is not a supported channel.", channel), http.StatusBadRequest))
	}
	input1, err := utils.StrToInt(getQueryParam(q, "input1"))
	if err != nil {
		panic(utils.NewCustomError("'input1' is a required integer parameter.", http.StatusBadRequest))
	}

	var args = make([]interface{}, 0)
	args = append(args, input1)
	var outputCache = make([][]interface{}, 0, len(channelConfig.StepsSequence))
	var currentOutput []interface{}
	for seqIndex, seq := range channelConfig.StepsSequence {
		// Reset args for preparation from second step onwards
		if seqIndex > 0 {
			args = nil
		}
		var value interface{}

		// Process input args as per step configuration
		for inputIndex, inputConfig := range seq.InputConfig {
			if inputIndex < len(args) {
				continue
			}
			argStrIndexes := strings.Split(inputConfig.Input, ".")

			if len(argStrIndexes) == 2 {
				// If current input is referred to any of previous step outputs, process from that
				argIntIndexes := make([]int, 0)
				for _, argStrValue := range argStrIndexes {
					outputIndex, err := utils.StrToInt(argStrValue)
					if err != nil {
						panic(utils.NewCustomError(fmt.Sprintf("Invalid step configuration for '%s'", seq.StepName), http.StatusBadRequest))
					}
					argIntIndexes = append(argIntIndexes, outputIndex)
				}
				value = outputCache[argIntIndexes[0]][argIntIndexes[1]]
			} else {
				// Pick default
				value, err = utils.ConvertStrToDatatype(inputConfig.Default, inputConfig.DataType)
				if err != nil {
					panic(utils.NewCustomError(fmt.Sprintf("Invalid default value for input {%d} for channel '%s'", inputIndex, channel), http.StatusBadRequest))
				}
			}
			args = append(args, value)
		}
		currentOutput, err = executePlugin(seq.StepName, args...)
		outputCache = append(outputCache, currentOutput)
		if err != nil {
			break
		}
	}

	if err != nil {
		panic(err)
	} else {
		c.String(http.StatusOK, "Uploading to %s response is %v...", c.Param("channel"), outputCache[len(channelConfig.StepsSequence)-1])
	}
}

func (s *ZServer) registerRoutes() {
	router := gin.Default()
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Accept", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/", defaultHandler)
	router.GET("/json", jsonHandler)

	router.GET("/channelConfig", getChannelConfig)
	router.OPTIONS("/channelConfig", defaultHandler)
	router.POST("/channelConfig", setChannelConfig)
	router.GET("/pluginsConfig", getPluginsConfig)

	router.GET("/channel/:channel", upload)
	s.server.Handler = router
}
