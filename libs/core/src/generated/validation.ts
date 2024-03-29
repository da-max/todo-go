import { z } from "zod";
import {
    ChangePassword,
    ConfirmIdentifier,
    Identifier,
    NewTodo,
    NewUser,
    RequestPasswordResetIdentifier,
    ResetPasswordIdentifier,
    UpdateUser,
} from "./types";

type Properties<T> = Required<{
    [K in keyof T]: z.ZodType<T[K], any, T[K]>;
}>;

type definedNonNullAny = {};

export const isDefinedNonNullAny = (v: any): v is definedNonNullAny =>
    v !== undefined && v !== null;

export const definedNonNullAnySchema = z
    .any()
    .refine((v) => isDefinedNonNullAny(v));

export function ChangePasswordSchema(): z.ZodObject<
    Properties<ChangePassword>
> {
    return z.object({
        confirmPassword: z
            .string({
                required_error: "Ce champs est requis.",
                invalid_type_error: "Ce champs est requis.",
            })
            .min(1),
        oldPassword: z
            .string({
                required_error: "Ce champs est requis.",
                invalid_type_error: "Ce champs est requis.",
            })
            .min(1),
        password: z
            .string({
                required_error: "Ce champs est requis.",
                invalid_type_error: "Ce champs est requis.",
            })
            .min(1),
    });
}

export function ConfirmIdentifierSchema(): z.ZodObject<
    Properties<ConfirmIdentifier>
> {
    return z.object({
        token: z
            .string({
                required_error: "Ce champs est requis.",
                invalid_type_error: "Ce champs est requis.",
            })
            .min(1),
    });
}

export function IdentifierSchema(): z.ZodObject<Properties<Identifier>> {
    return z.object({
        password: z
            .string({
                required_error: "Ce champs est requis.",
                invalid_type_error: "Ce champs est requis.",
            })
            .min(1),
        username: z
            .string({
                required_error: "Ce champs est requis.",
                invalid_type_error: "Ce champs est requis.",
            })
            .min(1),
    });
}

export function NewTodoSchema(): z.ZodObject<Properties<NewTodo>> {
    return z.object({
        text: z
            .string({
                required_error: "Ce champs est requis.",
                invalid_type_error: "Ce champs est requis.",
            })
            .min(1),
    });
}

export function NewUserSchema(): z.ZodObject<Properties<NewUser>> {
    return z.object({
        email: z
            .string({
                required_error: "Ce champs est requis.",
                invalid_type_error: "Ce champs est requis.",
            })
            .email()
            .min(5),
        password: z
            .string({
                required_error: "Ce champs est requis.",
                invalid_type_error: "Ce champs est requis.",
            })
            .min(1),
        username: z
            .string({
                required_error: "Ce champs est requis.",
                invalid_type_error: "Ce champs est requis.",
            })
            .min(1),
    });
}

export function RequestPasswordResetIdentifierSchema(): z.ZodObject<
    Properties<RequestPasswordResetIdentifier>
> {
    return z.object({
        email: z
            .string({
                required_error: "Ce champs est requis.",
                invalid_type_error: "Ce champs est requis.",
            })
            .email()
            .min(5),
    });
}

export function ResetPasswordIdentifierSchema(): z.ZodObject<
    Properties<ResetPasswordIdentifier>
> {
    return z.object({
        password: z
            .string({
                required_error: "Ce champs est requis.",
                invalid_type_error: "Ce champs est requis.",
            })
            .min(1),
        token: z
            .string({
                required_error: "Ce champs est requis.",
                invalid_type_error: "Ce champs est requis.",
            })
            .min(1),
    });
}

export function UpdateUserSchema(): z.ZodObject<Properties<UpdateUser>> {
    return z.object({
        email: z
            .string({
                required_error: "Ce champs est requis.",
                invalid_type_error: "Ce champs est requis.",
            })
            .email()
            .min(5)
            .nullish(),
        username: z
            .string({
                required_error: "Ce champs est requis.",
                invalid_type_error: "Ce champs est requis.",
            })
            .min(1)
            .nullish(),
    });
}
