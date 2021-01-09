import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';

import { Products, ProductsService } from '../core';

@Component({
  selector: 'app-editor-page',
  templateUrl: './editor.component.html'
})
export class EditorComponent implements OnInit {
  product: Products = {} as Products;
  productForm: FormGroup;
  errors: Object = {};
  isSubmitting = false;

  constructor(
    private productService: ProductsService,
    private route: ActivatedRoute,
    private router: Router,
    private fb: FormBuilder,
    private toastr: ToastrService
  ) {
    // use the FormBuilder to create a form group
    this.productForm = this.fb.group({
      Pname: '',
      Pbrand:'',
      description: '',
      Prating:'',
      Pcategory:''
    });
  }

  ngOnInit() {
    // If there's an product prefetched, load it
    this.route.data.subscribe((data: { product: Products }) => {
      if (data.product) {
        this.product = data.product;
        this.productForm.patchValue(data.product);
      }
    });
  }


  submitForm() {
    console.log('click')
    console.log(this.product);
    
    this.isSubmitting = true;

    // update the model
    this.updateProduct(this.productForm.value);

    // post the changes
    this.productService.save(this.product).subscribe(
      data => {
        console.log(data);
        this.toastr.success('Product Created', 'Create');
        this.router.navigateByUrl('/')
      }, 
      err => {
        this.toastr.error('Something Happens...', 'Create');
        this.errors = err;
        this.isSubmitting = false;
      }
    );
  }

  updateProduct(values: Object) {
    Object.assign(this.product, values);
  }
}
