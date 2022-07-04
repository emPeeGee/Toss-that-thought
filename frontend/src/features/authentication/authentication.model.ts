export interface CredentialsModel {
  username: string;
  password: string;
}

export interface AuthenticationResponse {
  token: string;
}

export interface UserModel {
  name: string;
  username: string;
  createdAt: string;
  updatedAt: string;
}
