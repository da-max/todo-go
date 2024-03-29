import gql from "graphql-tag";
import { userFragment } from "./fragments/user";

export const currentUser = gql`
    query currentUser {
        currentUser {
            ...User
        }
    }
    ${userFragment}
`;

export const loginMutation = gql`
    mutation login($input: Identifier!) {
        login(input: $input) {
            accessToken
            refreshToken
        }
    }
`;

export const signUp = gql`
    mutation signUp($input: NewUser!) {
        signUp(input: $input) {
            ...User
        }
    }
    ${userFragment}
`;

export const confirmAccount = gql`
    mutation confirmAccount($input: ConfirmIdentifier!) {
        confirmAccount(input: $input) {
            ok
        }
    }
`;

export const requestResetPassword = gql`
    mutation requestResetPassword($input: RequestPasswordResetIdentifier!) {
        requestResetPassword(input: $input) {
            ok
        }
    }
`;

export const requestConfirmAccount = gql`
    mutation requestConfirmAccount {
        requestConfirmAccount {
            ok
        }
    }
`;

export const resetPassword = gql`
    mutation resetPassword($input: ResetPasswordIdentifier!) {
        resetPassword(input: $input) {
            ok
        }
    }
`;

export const updateAccount = gql`
    mutation updateAccount($input: UpdateUser!) {
        updateAccount(input: $input) {
            ...User
        }
    }
    ${userFragment}
`;

export const deleteAccount = gql`
    mutation deleteAccount {
        deleteAccount {
            ok
        }
    }
`;

export const changePassword = gql`
    mutation changePassword($input: ChangePassword!) {
        changePassword(input: $input) {
            ok
        }
    }
`;
