export interface ThoughtMetadataModel {
  thoughtKey?: string;
  metadataKey: string;
  lifetime: string;
  passphrase?: string;
  abbreviatedThoughtKey?: string;
}

export interface ThoughtCreateRequest {
  thought: string;
  passphrase?: string;
  lifetime: string;
}
