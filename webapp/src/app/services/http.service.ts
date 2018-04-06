import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';

import { ChannelService } from './channel.service';

@Injectable()
export class HttpService {
  constructor(public http: HttpClient) { }

  public getHeader(tokenPrefix = '') {
    const headers = new HttpHeaders()
      .append('Content-Type', 'application/json')
      .append('Accept', 'application/json');
    return { headers };
  }

  public doGet(url, header = this.getHeader()) {
    return this.http.get(url, header).toPromise();
  }

  public doPost(url, json, header = this.getHeader()) {
    return this.http.post(url, json, header).toPromise();
  }

  public doPut(url, json, header = this.getHeader()) {
    return this.http.put(url, json, header).toPromise();
  }

  public doDelete(url, header = this.getHeader()) {
    return this.http.delete(url, header).toPromise();
  }

  public doPatch(url, json, header = this.getHeader()) {
    return this.http.patch(url, json, header).toPromise();
  }
}
