import { ExceptionFilter, Catch, ArgumentsHost } from '@nestjs/common';
import { Response } from 'express';
import { PixKeyGrpcUnknownError } from '../pix-keys.services';

@Catch(PixKeyGrpcUnknownError)
export class PixKeyGrpcUnknownErrorFilter implements ExceptionFilter {
  catch(exception: PixKeyGrpcUnknownError, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();

    response.status(500).json({
      statusCode: 500,
      message: exception.message,
    });
  }
}
