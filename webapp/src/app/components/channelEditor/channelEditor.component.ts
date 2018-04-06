import { Component, OnInit, OnChanges, Input } from '@angular/core';

@Component({
  selector: 'app-channel-editor',
  templateUrl: './channelEditor.component.html',
  styleUrls: ['./channelEditor.component.scss']
})
export class ChannelEditorComponent implements OnInit, OnChanges {
  @Input() channelName = '';
  @Input() systemConfig: any = {};
  @Input() pluginsConfig: any = {};
  @Input() updateChannelName: any;

  public dragging = false;
  public dragStageDropIndex = -1;
  private dragStageIndex = 0;
  private _channelName = '';

  constructor() {
    this.dropEnd = this.dropEnd.bind(this);
    this.dragStage = this.dragStage.bind(this);
    this.dropStage = this.dropStage.bind(this);
    this.deleteStep = this.deleteStep.bind(this);
    this.getDupStepConfig = this.getDupStepConfig.bind(this);
  }

  public ngOnInit() {
  }

  public ngOnChanges() {
    this._channelName = this.channelName;
  }

  public _updateChannelName() {
    if (this.updateChannelName(this.channelName, this._channelName)) {
      this.channelName = this._channelName;
      this.systemConfig.Channels[this.channelName].$modified = true;
    } else {
      this._channelName = this.channelName;
    }
  }

  public addStep(index = this.systemConfig.Channels[this.channelName].StepsSequence.length) {
    let newStepName;
    this.systemConfig.Channels[this.channelName].$modified = true;
    if (index === 0) {
      newStepName = Object.keys(this.pluginsConfig).find(plugin => this.pluginsConfig[plugin].CanBeFirst);
    } else {
      newStepName = Object.keys(this.pluginsConfig).find(plugin => !this.pluginsConfig[plugin].ShouldBeFirstOnly);
    }

    this.systemConfig.Channels[this.channelName].StepsSequence.splice(index, 0, this.getDupStepConfig(newStepName));
  }

  public deleteStep(index) {
    if (this.systemConfig.Channels[this.channelName].StepsSequence.length > 1) {
      this.systemConfig.Channels[this.channelName].StepsSequence.splice(index, 1);
      this.systemConfig.Channels[this.channelName].$modified = true;
    }
  }

  public getDupStepConfig(stepName) {
    const dupStep = JSON.parse(JSON.stringify(this.pluginsConfig[stepName]));
    dupStep.InputConfig.forEach(ic => ic.Input = '');
    const newStep = {
      StepName: stepName,
      InputConfig: dupStep.InputConfig
    };

    return newStep;
  }

  public dragStage(e, index) {
    e.dataTransfer.effectAllowed = 'move';
    this.dragging = true;
    this.dragStageIndex = index;
  }

  public dragoverStage(e, index, activate) {
    e.preventDefault();
    e.dataTransfer.dropEffect = 'move';
    this.dragging = activate && !((this.dragStageIndex === index) || (this.dragStageIndex === index - 1));
    this.dragStageDropIndex = this.dragging ? index : -1;
  }

  public dropStage(e, targetIndex) {
    e.preventDefault();
    if ((this.dragStageIndex === targetIndex) || (this.dragStageIndex === targetIndex - 1)) {
      // Take no action
      return;
    }

    if (this.dragStageIndex < targetIndex) {
      targetIndex--;
    }
    const sourceConfig = this.systemConfig.Channels[this.channelName].StepsSequence[this.dragStageIndex];
    this.systemConfig.Channels[this.channelName].StepsSequence.splice(this.dragStageIndex, 1);
    this.systemConfig.Channels[this.channelName].StepsSequence.splice(targetIndex, 0, sourceConfig);
    this.systemConfig.Channels[this.channelName].$modified = true;
  }

  public dropEnd() {
    this.dragging = false;
  }
}
