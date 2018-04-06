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
    this.refreshPluginsConfig();
  }

  public refreshChannelsConfig() {
    this.httpService.doGet(`${environment.uploaderEndPoint}/channelConfig`)
      .then((res) => {
        this.channelsConfig = res;
        this.channelsConfigOriginal = JSON.parse(JSON.stringify(res));
      });
  }

  public refreshPluginsConfig() {
    this.httpService.doGet(`${environment.uploaderEndPoint}/pluginsConfig`)
      .then((res) => {
        this.pluginsConfig = res;
      });
  }

  public updateChannelConfig(channelName, channelConfig) {
    const payload = {
      ChannelName: channelName,
      ChannelConfig: channelConfig
    };

    console.log(payload);

    this.httpService.doPost(`${environment.uploaderEndPoint}/channelConfig`, payload)
      .then((res) => {
        return this.refreshChannelsConfig();
      });
  }
}
