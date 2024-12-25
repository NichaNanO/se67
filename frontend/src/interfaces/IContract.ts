import { ContractStatus } from "./IContractStatus";
import { ContractType } from "./IContractType";
import { Employee } from "./IEmployee";
import { Member } from "./IMember";

export interface Contract {
    id: number;
    startDate: string; // in ISO format
    endDate: string; // in ISO format
    securityDeposit: number;
    note: string;
    memberId: number;
    employeeId: number;
    roomId: number;
    contractTypeId: number;
    statusId: number;
    member: Member;
    employee: Employee;
    contractType: ContractType;
    status: ContractStatus;
  }
  