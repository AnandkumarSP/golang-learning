import { ModuleWithProviders } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { ChannelsComponent } from './containers/channels/channels.component';

export const appRoutes: Routes = [
  { path: '', redirectTo: 'channels', pathMatch: 'full' },
  { path: 'channels', component: ChannelsComponent, pathMatch: 'full' },
  { path: '**', redirectTo: '/channels', pathMatch: 'full' }
];

