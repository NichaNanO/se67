import { Contract } from "./IContract";

export interface Room {
    id: number;
    roomNumber: string;
    roomType: string;
    price: number;
    contracts: Contract[];
  }
  