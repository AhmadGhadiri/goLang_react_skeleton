import Posts from '../components/Posts/Posts';
import AuthContext from '../store/auth-context';
import { useContext } from 'react';
import { Navigate } from 'react-router-dom';

const PostsPage = () => {
    const authContext = useContext(AuthContext);
    if (authContext.loggedIn) {

        return (<Posts />);
    }
    else {
        return (<Navigate replace to="/auth" />);
    }
};

export default PostsPage;