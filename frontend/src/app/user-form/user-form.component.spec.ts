import { ComponentFixture, TestBed } from '@angular/core/testing';
import { UserFormComponent } from './user-form.component';
import { UserService } from '../user.service';
import { ReactiveFormsModule } from '@angular/forms';
import { MatCardModule } from '@angular/material/card';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSelectModule } from '@angular/material/select';
import { MatInputModule } from '@angular/material/input';
import { RouterTestingModule } from '@angular/router/testing';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { of } from 'rxjs';

describe('UserFormComponent', () => {
  let component: UserFormComponent;
  let fixture: ComponentFixture<UserFormComponent>;
  let userServiceSpy: jasmine.SpyObj<UserService>;

  beforeEach(async () => {
    const spy = jasmine.createSpyObj('UserService', [
      'createUser',
      'updateUser',
    ]);

    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule,
        MatCardModule,
        MatSnackBarModule,
        MatFormFieldModule,
        MatSelectModule,
        MatInputModule,
        RouterTestingModule,
        BrowserAnimationsModule,
      ],
      declarations: [UserFormComponent],
      providers: [{ provide: UserService, useValue: spy }],
    }).compileComponents();

    userServiceSpy = TestBed.inject(UserService) as jasmine.SpyObj<UserService>;
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(UserFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should initialize the form', () => {
    expect(component.userForm).toBeDefined();
    expect(component.userForm.contains('user_name')).toBeTruthy();
    expect(component.userForm.contains('first_name')).toBeTruthy();
    expect(component.userForm.contains('last_name')).toBeTruthy();
    expect(component.userForm.contains('email')).toBeTruthy();
    expect(component.userForm.contains('user_status')).toBeTruthy();
    expect(component.userForm.contains('department')).toBeTruthy();
  });

  it('should call createUser when submitting a new user', () => {
    const testUser = {
      user_name: 'test',
      first_name: 'Test',
      last_name: 'User',
      email: 'test@example.com',
      user_status: 'A',
      department: 'IT',
    };
    component.userForm.setValue(testUser);
    userServiceSpy.createUser.and.returnValue(of({ id: 1, ...testUser }));

    component.onSubmit();

    expect(userServiceSpy.createUser).toHaveBeenCalledWith(testUser);
  });
});
