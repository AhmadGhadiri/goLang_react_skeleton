import { useContext } from 'react';
import { Link, useNavigate } from 'react-router-dom';

import AuthContext from '../../store/auth-context';

const NavigationBar = () => {
    const navigate = useNavigate();

    const authContext = useContext(AuthContext);

    const loggedIn = authContext.loggedIn;
    const currentUser = authContext.currentUser;

    const logoutHandler = () => {
        authContext.logout();
        navigate('/');
    };

    return (
        <nav className="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
            <button type="button" className="navbar-toggler" data-toggle="collapse" data-target="#navbarCollapse"
                aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
                <span className="navbar-toggler-icon"></span>
            </button>
            <div className="collapse navbar-collapse justify-content-between" id="navbarCollapse">
                <div className="d-flex flex-row">
                    <ul className="navbar-nav mr-auto">
                        <li key="uniq1" className="nav-item">
                            <Link className="nav-link" to="/">RGB</Link>
                        </li>
                        {loggedIn && (
                            <li key="uniq2" className="nav-item">
                                <Link className="nav-link" to="/posts">{currentUser}'s posts</Link>
                            </li>
                        )}
                    </ul>
                </div>
                <div className="d-flex flex-row justify-content-end">
                    <ul className="navbar-nav mr-auto">
                        {!loggedIn && (
                            <li key="uniq3" className="nav-item">
                                <Link className="nav-link" to="/auth">Login</Link>
                            </li>
                        )}
                        {loggedIn && (
                            <li key="uniq4" className="nav-item">
                                <button className="btn btn-dark" onClick={logoutHandler}>LOGOUT</button>
                            </li>
                        )}
                    </ul>
                </div>
            </div>
        </nav>
    );
}

export default NavigationBar;