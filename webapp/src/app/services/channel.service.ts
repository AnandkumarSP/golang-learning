import { Injectable } from '@angular/core';

import { environment } from './../../environments/environment';
import { HttpService } from './http.service';

@Injectable()
export class ChannelService {
  public channelsConfig: any = {};
  public channelsConfigOriginal: any = {};
  public pluginsConfig: any = {};

  constructor(public httpService: HttpService) {
    this.init();
  }

  public init() {
    this.refreshChannelsConfig();
    this.refreshSystemConfig();
    this.refreshPluginsConfig();
  }

  public refreshChannelsConfig() {
    this.httpService.doGet(`${environment.uploaderEndPoint}/workflows`)
      .then((res) => {
        this.channelsConfig.Channels = res;
        this.channelsConfigOriginal.Channels = JSON.parse(JSON.stringify(res));
      });
  }

  public refreshSystemConfig() {
    this.httpService.doGet(`${environment.uploaderEndPoint}/systemconfig`)
      .then((res) => {
        this.channelsConfig.FirstStepInputConfig = res['inputs'];
      });
  }

  public refreshPluginsConfig() {
    this.httpService.doGet(`${environment.uploaderEndPoint}/plugins`)
      .then((res) => {
        this.pluginsConfig = res;
      });
  }

  public updateChannelConfig(channelName, channelConfig) {
    this.httpService.doPost(`${environment.uploaderEndPoint}/workflow/${channelName}`, channelConfig)
      .then((res) => {
        return this.refreshChannelsConfig();
      });
  }
}
