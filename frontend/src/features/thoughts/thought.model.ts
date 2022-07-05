export interface ThoughtMetadataModel {
  thoughtKey?: string;
  metadataKey: string;
  lifetime: string;
  passphrase?: string;
  abbreviatedThoughtKey?: string;
  burnedDate?: string;
  viewedDate?: string;
  isBurned: boolean;
  isViewed: boolean;
  createdDate?: string;
}

export interface ThoughtCreateRequest {
  thought: string;
  passphrase?: string;
  lifetime: string;
}

export interface ThoughtPassphraseRequest {
  passphrase: string;
}

export interface ThoughtResponse {
  thought: string;
}

export interface ThoughtPassphraseInfo {
  canPassphraseBeSkipped: boolean;
}
