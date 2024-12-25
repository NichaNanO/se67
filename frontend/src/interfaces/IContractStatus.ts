import { Contract } from "./IContract";

export interface ContractStatus {
    id: number;
    status: string;
    contracts: Contract[];
  }
  