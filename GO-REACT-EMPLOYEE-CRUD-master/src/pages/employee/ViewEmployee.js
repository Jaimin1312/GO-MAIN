import React, { useState, useEffect } from "react";
import axios from 'axios';
import { useHistory, useParams ,Link} from 'react-router-dom'
function ViewEmployee() {
    const { id } = useParams();
    let history = useHistory();
    const [Employee, setEmployee] = useState({
        name: "",
        city: "",
        mobile: "",
        email: "",
    });
    const { name, city, mobile, email } = Employee;
    const loadEmployee = async () => {
        const result = await axios.get(`http://localhost:9000/api/employees/${id}`)
        setEmployee(result.data)
    }

    useEffect(() => {
        loadEmployee();
    }, []);


    return (
        <div className="container py-4">
            <Link className="btn btn-primary" to="/">
                back to Home
        </Link>
            <h1 className="display-4">Employee Id: {id}</h1>
            <hr />
            <ul className="list-group w-50">
                <li className="list-group-item">name: {Employee.name}</li>
                <li className="list-group-item">city: {Employee.city}</li>
                <li className="list-group-item">email: {Employee.email}</li>
                <li className="list-group-item">phone: {Employee.mobile}</li>
            </ul>
        </div>
    );
}

export default ViewEmployee;