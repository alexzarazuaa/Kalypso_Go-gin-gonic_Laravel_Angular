import { NgModule } from '@angular/core';

import { EditorComponent } from './editor.component';
import { EditableProduct } from './editable-product-resolver.service';
import { SharedModule } from '../shared';
import { EditorRoutingModule } from './editor-routing.module';

@NgModule({
  imports: [SharedModule, EditorRoutingModule],
  declarations: [EditorComponent],
  providers: [EditableProduct]
})
export class EditorModule {}
