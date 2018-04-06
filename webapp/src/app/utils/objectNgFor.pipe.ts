import { Pipe, PipeTransform } from '@angular/core';

@Pipe({ name: 'ObjNgFor', pure: false })
export class ObjNgFor implements PipeTransform {
  transform(value: any, sort = false, pickValue = false): any {
    let retval = [];

    if (value && typeof value === 'object') {
      if (pickValue) {
        retval = Object.values(value);
      } else {
        retval = Object.keys(value);
        if (sort) {
          retval = retval.sort();
        }
      }
      return retval;
    }
  }
}
