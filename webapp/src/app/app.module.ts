import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';

import { appRoutes } from './app.routing';
import { AppComponent } from './app.component';
import { ChannelEditorComponent, StepEditorComponent } from './components';
import { ChannelsComponent } from './containers';
import { ChannelService, HttpService } from './services';
import { ObjNgFor } from './utils';

@NgModule({
  declarations: [
    AppComponent,
    ChannelsComponent,
    ChannelEditorComponent,
    StepEditorComponent,
    ObjNgFor
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    FormsModule,
    RouterModule.forRoot(appRoutes)
  ],
  providers: [
    ChannelService,
    HttpService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
