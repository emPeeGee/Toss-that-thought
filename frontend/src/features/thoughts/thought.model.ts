// TODO: Here something is not correct, replace NullTime on back
interface NullTime {
  Time: string;
  Valid: boolean;
}

export interface ThoughtMetadataModel {
  thoughtKey?: string;
  metadataKey: string;
  lifetime: string;
  passphrase?: string;
  abbreviatedThoughtKey?: string;
  burnedDate?: NullTime;
  viewedDate?: NullTime;
  isBurned: boolean;
  isViewed: boolean;
}

export interface ThoughtCreateRequest {
  thought: string;
  passphrase?: string;
  lifetime: string;
}

export interface ThoughtBurnRequest {
  passphrase: string;
}
