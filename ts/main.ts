import './auth';
import {Amplify} from 'aws-amplify';

if (window.appConfig) {
    Amplify.configure({
        Auth: {
            Cognito: {
                userPoolId: appConfig.userPoolId,
                userPoolClientId: appConfig.userPoolWebClientId
            }
        }
    })
}
