import { Component, OnInit, Input } from '@angular/core';

const noop = function () {};

@Component({
  selector: 'app-step-editor',
  templateUrl: './stepEditor.component.html',
  styleUrls: ['./stepEditor.component.scss']
})
export class StepEditorComponent implements OnInit {
  @Input() index = 1;
  @Input() stepConfig: any = {};
  @Input() channelName = '';
  @Input() systemConfig: any = {};
  @Input() pluginsConfig: any = {};
  @Input() getDupStepConfig: any = noop;
  @Input() deleteStep: any = noop;
  @Input() dragStage: any = noop;
  @Input() dropEnd: any = noop;

  public showDesc = false;

  public ngOnInit() {
    console.log(this);
  }

  public toggleInputDescription() {
    this.showDesc = !this.showDesc;
  }

  public channelModified() {
    this.systemConfig.Channels[this.channelName].$modified = true;
  }

  public stepChanged() {
    const newStepConfig = this.getDupStepConfig(this.stepConfig.StepName);

    Object.keys(this.stepConfig).forEach(k => delete this.stepConfig[k]);
    Object.keys(newStepConfig).forEach(k => this.stepConfig[k] = newStepConfig[k]);
    this.channelModified();
  }

  public getPossibleInputs(index, inputConfig) {
    const retval = [{
      name: 'None',
      value: ''
    }];

    this.systemConfig['FirstStepInputConfig']
      .forEach((sc) => {
        if (sc.DataType === inputConfig.DataType) {
          retval.push({
            name: `${sc.Name}`,
            value: `G${sc.Name}`
          });
        }
      });

    this.systemConfig.Channels[this.channelName]['StepsSequence']
      .filter((ss, stepIndex) => stepIndex < index)
      .map(ss => ss.StepName)
      .forEach((sn, stepIndex) => {
        this.pluginsConfig[sn]['OutputConfig']
          .map((oc, inputIndex) => {
            if (oc.DataType === inputConfig.DataType) {
              retval.push({
                name: `Step${stepIndex + 1} (${sn}): ${oc.Name}`,
                value: `${stepIndex}.${inputIndex}`
              });
            }
          });
      });

    if (!retval.find(c => c.value === inputConfig.Input)) {
      setTimeout(() => {
        inputConfig.Input = '';
      });
    }

    return retval;
  }
}
