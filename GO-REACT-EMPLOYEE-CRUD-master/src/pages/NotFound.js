
import React from 'react'
import { Link, NavLink } from 'react-router-dom'
function NotFound() {
    return (
        <div className="not-found">
            <div className="jumbotron text-center">
                <h1 className="display-4">PAGE NOT FOUND</h1>
                <h4>Sorry your requested page does not exits or missing</h4>
                <div className="row">
                    <div className="col text-center">
                        <Link className="nav-link" exact to="/">
                            <button type="button" className="btn btn-primary">
                                Home
                         </button>
                        </Link>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default NotFound;
