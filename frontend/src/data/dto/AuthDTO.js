import { User } from '@/domain/entities/User';

export class AuthDTO {
  static toEntity(data) {
    return new User({
      email: data.email,
      role: data.role,
      token: data.token,
    });
  }
}
