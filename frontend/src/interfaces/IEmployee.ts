import { Gender } from "./IGenders";
import { Member } from "./IMember";
import { Position } from "./IPosition";

export interface Employee {
    id: number;
    firstName: string;
    lastName: string;
    phoneNumber: string;
    nationalId: string;
    email: string;
    password: string;
    profile: string;
    genderId: number;
    gender: Gender;
    positionId: number;
    position: Position;
    members: Member[];
  }
  