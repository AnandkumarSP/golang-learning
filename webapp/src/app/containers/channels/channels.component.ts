// tslint:disable:max-line-length
import { Component } from '@angular/core';

import { ChannelService } from './../../services';

@Component({
  selector: 'app-channels',
  templateUrl: './channels.component.html',
  styleUrls: ['./channels.component.scss']
})
export class ChannelsComponent {
  public selectedChannel = '';

  constructor(public channelService: ChannelService) {
    console.log(this);
    this.updateChannelName = this.updateChannelName.bind(this);
  }

  public selectChannel(channel: string) {
    this.selectedChannel = channel;
  }

  private getUniqueChannelName() {
    let i = 1;
    while (i) {
      if (!this.channelService.channelsConfig.Channels[`<NewChannel${i}>`]) {
        return `<NewChannel${i}>`;
      }
      i++;
    }
  }

  public addChannel() {
    this.channelService.channelsConfig.Channels[this.getUniqueChannelName()] = {
      StepsSequence: [],
      $new: true
    };
  }

  public updateChannelName(oldName, newName) {
    if (this.channelService.channelsConfig.Channels[newName]) {
      return;
    } else {
      this.channelService.channelsConfig.Channels[newName] = this.channelService.channelsConfig.Channels[oldName];
      delete this.channelService.channelsConfig.Channels[oldName];
      if (oldName === this.selectedChannel) {
        this.selectedChannel = newName;
      }
      return true;
    }
  }

  public saveChannel() {
    const channelConfig = JSON.parse(JSON.stringify(this.channelService.channelsConfig.Channels[this.selectedChannel], (k, v) => k.startsWith('$') ? undefined : v));
    this.channelService.updateChannelConfig(this.selectedChannel, channelConfig);
  }

  public revertChanges() {
    this.channelService.channelsConfig.Channels[this.selectedChannel] = JSON.parse(JSON.stringify(this.channelService.channelsConfigOriginal.Channels[this.selectedChannel]));
  }
}
