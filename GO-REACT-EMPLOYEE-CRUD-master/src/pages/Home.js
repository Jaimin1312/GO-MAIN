
import React, { useState, useEffect } from 'react'
import axios from 'axios'
import {Link} from 'react-router-dom'
function Home() {
    const [Employees, setEmployee] = useState([]);

    const loadEmployee = async () => {
        const result = await axios.get("http://localhost:9000/api/employees");
        //console.log(result.data)
        setEmployee(result.data)
    }

    useEffect(() => {
        loadEmployee()
    }, [])

    const deleteEmployee = async id => {
        console.log(id)
        await axios.delete(`http://localhost:9000/api/employees/${id}`);
        loadEmployee();
    };

    return (
        <div className="container">
            <div className="py-4 font-weight-bold">
                <table class="table border shadow">
                    <thead class="thead-dark">
                        <tr>
                            <th scope="col">#</th>
                            <th scope="col">Name</th>
                            <th scope="col">City</th>
                            <th scope="col">mobile</th>
                            <th scope="col">email</th>
                            <th scope="col">Action</th>
                        </tr>
                    </thead>
                    <tbody>
                        {
                            Employees.map((employee, index) => (
                                <tr scope="row">
                                    <td> {employee.id} </td>
                                    <td> {employee.name} </td>
                                    <td> {employee.city} </td>
                                    <td> {employee.mobile} </td>
                                    <td> {employee.email} </td>
                                    <td>
                                        <Link className="btn btn-primary mr-5" exact to={`/employee/${employee.id}`}>
                                            View
                                        </Link>
                                        <Link className="btn btn-info mr-5" exact to={`/employee/edit/${employee.id}`}>
                                            Edit
                                        </Link>
                                        <Link className="btn btn-danger" onClick={() => deleteEmployee(employee.id)}>
                                            Delete
                                        </Link>
                                    </td>
                                </tr>
                            ))
                        }
                    </tbody>
                </table>
            </div>
        </div>
    );
}

export default Home;
