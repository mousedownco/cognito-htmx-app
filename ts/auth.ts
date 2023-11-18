import {
    autoSignIn,
    confirmSignUp,
    type ConfirmSignUpInput,
    signIn,
    type SignInInput,
    signOut,
    signUp
} from 'aws-amplify/auth';

/**
 * @file auth.ts
 * @see https://docs.amplify.aws/javascript/build-a-backend/auth/enable-sign-up/
 */

type SignUpParams = {
    username: string;
    password: string;
    email: string;
    phone_number?: string;
}

async function handleSignUp({username, password, email, phone_number}: SignUpParams) {
    try {
        const {isSignUpComplete, userId, nextStep} = await signUp({
            username,
            password,
            options: {
                userAttributes: {
                    email,
                    phone_number
                },
                autoSignIn: true
            }
        });
        console.log('userId:', userId);
    } catch (error) {
        console.log('error signing up:', error);
    }
}

async function handleSignUpConfirmation({username, confirmationCode}: ConfirmSignUpInput) {
    try {
        const {isSignUpComplete, nextStep} = await confirmSignUp({username, confirmationCode});
        console.log('isSignUpComplete:', isSignUpComplete);
    } catch (error) {
        console.log('error confirming sign up:', error);
    }
}

async function handleAutoSignIn() {
    try {
        const user = await autoSignIn();
        console.log('user:', user);
    } catch (error) {
        console.log('error signing in:', error);
    }
}

async function handleSignIn({username, password}: SignInInput) {
    try {
        const {isSignedIn, nextStep} = await signIn({username, password});
        console.log('nestStep:', nextStep);
    } catch (error) {
        console.log('error signing in:', error);
    }
}

async function handleSignOut() {
    try {
        await signOut();
    } catch (error) {
        console.log('error signing out:', error);
    }
}