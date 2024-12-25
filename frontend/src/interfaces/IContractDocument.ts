import { Contract } from "./IContract";

export interface ContractDocument {
    id: number;
    contractId: number;
    documentType: string; // e.g., 'signature', 'file'
    filePath: string; // or the base64 encoded string for signatures
    contract: Contract;
  }
  