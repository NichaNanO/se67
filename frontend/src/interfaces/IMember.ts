import { Employee } from "./IEmployee";
import { Gender } from "./IGenders";
import { Position } from "./IPosition";

export interface Member {
    id: number;
    firstName: string;
    lastName: string;
    phoneNumber: string;
    nationalId: string;
    email: string;
    password: string;
    profile: string;
    employeeId: number;
    employee: Employee;
    genderId: number;
    gender: Gender;
    positionId: number;
    position: Position;
  }
  