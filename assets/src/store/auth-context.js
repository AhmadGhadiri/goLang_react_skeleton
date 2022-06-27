import React, { useState } from 'react';

const AuthContext = React.createContext({
    token: null,
    loggedIn: false,
    login: (token) => { },
    logout: () => { },
});

export const AuthContextProvider = (props) => {
    const tokenKey = 'rgbToken';

    const [token, setToken] = useState(localStorage.getItem(tokenKey));
    const [currentuser, setCurrentUser] = useState(localStorage.getItem("currentUser"));


    const loggedIn = !!token;

    const loginHandler = (token, username) => {
        setToken(token);
        setCurrentUser(username);
        localStorage.setItem(tokenKey, token);
        localStorage.setItem("currentUser", username);
    };

    const logoutHandler = () => {
        setToken(null);
        setCurrentUser(null);
        localStorage.removeItem(tokenKey);
        localStorage.removeItem("currentUser");
    };

    const contextValue = {
        token: token,
        currentUser: currentuser,
        loggedIn: loggedIn,
        login: loginHandler,
        logout: logoutHandler,
    };

    return (
        <AuthContext.Provider value={contextValue}>
            {props.children}
        </AuthContext.Provider>
    );
};

export default AuthContext;