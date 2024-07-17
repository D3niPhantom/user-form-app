import { ComponentFixture, TestBed } from '@angular/core/testing';
import { UserListComponent } from './user-list.component';
import { UserService } from '../user.service';
import { of } from 'rxjs';
import { MatTableModule } from '@angular/material/table';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { RouterTestingModule } from '@angular/router/testing';
import { User } from '../models/user.model';

describe('UserListComponent', () => {
  let component: UserListComponent;
  let fixture: ComponentFixture<UserListComponent>;
  let userServiceSpy: jasmine.SpyObj<UserService>;

  beforeEach(async () => {
    const spy = jasmine.createSpyObj('UserService', ['getUsers', 'deleteUser']);

    await TestBed.configureTestingModule({
      imports: [
        MatCardModule,
        MatTableModule,
        MatIconModule,
        MatButtonModule,
        RouterTestingModule,
        BrowserAnimationsModule,
      ],
      declarations: [UserListComponent],
      providers: [{ provide: UserService, useValue: spy }],
    }).compileComponents();

    userServiceSpy = TestBed.inject(UserService) as jasmine.SpyObj<UserService>;
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(UserListComponent);
    component = fixture.componentInstance;
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should load users on init', () => {
    const testUsers: User[] = [
      {
        id: 1,
        user_name: 'test',
        first_name: 'Test',
        last_name: 'User',
        email: 'test@example.com',
        user_status: 'A',
        department: 'IT',
      },
    ];
    userServiceSpy.getUsers.and.returnValue(of(testUsers));

    fixture.detectChanges();

    expect(userServiceSpy.getUsers).toHaveBeenCalled();
    expect(component.dataSource.data).toEqual(testUsers);
  });
});
