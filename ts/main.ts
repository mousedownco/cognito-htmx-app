import {Amplify} from 'aws-amplify';
import {
    autoSignIn,
    confirmSignUp,
    type ConfirmSignUpInput,
    signIn,
    type SignInInput,
    signOut,
    signUp
} from 'aws-amplify/auth';

// if (window.appConfig) {
    console.log('configuring Amplify...');
    Amplify.configure({
        Auth: {
            Cognito: {
                userPoolId: appConfig.userPoolId,
                userPoolClientId: appConfig.userPoolWebClientId
            }
        }
    });
// }

export function configAuth(appConf: any) {
    console.log('configuring Amplify...');
    Amplify.configure({
        Auth: {
            Cognito: {
                userPoolId: appConf.userPoolId,
                userPoolClientId: appConf.userPoolWebClientId
            }
        }
    })
}

/**
 * @see https://docs.amplify.aws/javascript/build-a-backend/auth/enable-sign-up/
 */
type SignUpParams = {
    email: string;
    password: string;
}

export async function handleSignUp({email, password}: SignUpParams) {
    try {
        const {isSignUpComplete, userId, nextStep} = await signUp({
            username: email,
            password,
            options: {
                userAttributes: {
                    email,
                },
                autoSignIn: true
            }
        });
        console.log('userId:', userId);
    } catch (error) {
        console.log('error signing up:', error);
    }
}

export async function handleSignUpConfirmation({username, confirmationCode}: ConfirmSignUpInput) {
    try {
        const {isSignUpComplete, nextStep} = await confirmSignUp({username, confirmationCode});
        console.log('isSignUpComplete:', isSignUpComplete);
    } catch (error) {
        console.log('error confirming sign up:', error);
    }
}

export async function handleAutoSignIn() {
    try {
        const user = await autoSignIn();
        console.log('user:', user);
    } catch (error) {
        console.log('error signing in:', error);
    }
}

export async function handleSignIn({username, password}: SignInInput) {
    try {
        const {isSignedIn, nextStep} = await signIn({username, password});
        console.log('nestStep:', nextStep);
    } catch (error) {
        console.log('error signing in:', error);
    }
}

export async function handleSignOut() {
    try {
        await signOut();
    } catch (error) {
        console.log('error signing out:', error);
    }
}

(window as any).handleSignUp = handleSignUp;
(window as any).configAuth = configAuth;