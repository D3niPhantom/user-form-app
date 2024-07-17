import { TestBed } from '@angular/core/testing';
import {
  provideHttpClientTesting,
  HttpTestingController,
} from '@angular/common/http/testing';
import { provideHttpClient, HttpErrorResponse } from '@angular/common/http';
import { UserService } from './user.service';
import { MatSnackBar } from '@angular/material/snack-bar';
import { User } from './models/user.model';

describe('UserService', () => {
  let service: UserService;
  let httpMock: HttpTestingController;
  let snackBarSpy: jasmine.SpyObj<MatSnackBar>;

  beforeEach(() => {
    const spy = jasmine.createSpyObj('MatSnackBar', ['open']);

    TestBed.configureTestingModule({
      providers: [
        UserService,
        provideHttpClient(),
        provideHttpClientTesting(),
        { provide: MatSnackBar, useValue: spy },
      ],
    });
    service = TestBed.inject(UserService);
    httpMock = TestBed.inject(HttpTestingController);
    snackBarSpy = TestBed.inject(MatSnackBar) as jasmine.SpyObj<MatSnackBar>;
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should get users', () => {
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

    service.getUsers().subscribe((users) => {
      expect(users).toEqual(testUsers);
    });

    const req = httpMock.expectOne('http://localhost:8080/api/users');
    expect(req.request.method).toBe('GET');
    req.flush(testUsers);
  });

  it('should create a user', () => {
    const testUser: User = {
      user_name: 'test',
      first_name: 'Test',
      last_name: 'User',
      email: 'test@example.com',
      user_status: 'A',
      department: 'IT',
    };

    service.createUser(testUser).subscribe((user) => {
      expect(user).toEqual({ id: 1, ...testUser });
    });

    const req = httpMock.expectOne('http://localhost:8080/api/users');
    expect(req.request.method).toBe('POST');
    req.flush({ id: 1, ...testUser });
    expect(snackBarSpy.open).toHaveBeenCalledWith(
      'User created successfully',
      'Close',
      jasmine.any(Object)
    );
  });

  it('should handle errors when creating a user', () => {
    const testUser: User = {
      user_name: 'test',
      first_name: 'Test',
      last_name: 'User',
      email: 'test@example.com',
      user_status: 'A',
      department: 'IT',
    };

    service.createUser(testUser).subscribe({
      next: () => fail('should have failed with 404 error'),
      error: (error: HttpErrorResponse) => {
        expect(error.status).toBe(404);
        expect(error.error).toBe('User not found');
      },
    });

    const req = httpMock.expectOne('http://localhost:8080/api/users');
    expect(req.request.method).toBe('POST');
    req.flush('User not found', { status: 404, statusText: 'Not Found' });
  });
});
