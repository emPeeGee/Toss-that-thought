export interface ThoughtMetadataModel {
  thoughtKey?: string;
  metadataKey: string;
  lifetime: string;
  passphrase?: string;
}

export interface ThoughtCreateRequest {
  thought: string;
  passphrase?: string;
  lifetime: string;
}
