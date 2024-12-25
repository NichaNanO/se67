import { Contract } from "./IContract";

export interface ContractType {
    id: number;
    contractName: string;
    monthlyRent: number;
    durationMonths: number;
    contracts: Contract[];
  }
  