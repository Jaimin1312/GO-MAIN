import axios from 'axios';
import React,{useState} from 'react'
import {useHistory} from 'react-router-dom'
function AddEmployee() {
    let history = useHistory();
    const [Employee,setEmployee] = useState({
        name: "",
        city: "",
        mobile:"",
        email:"",
    });
    const { name, city, mobile, email } = Employee;
    const onInputChange = e => {
        setEmployee({ ...Employee,[e.target.name]: e.target.value})
    };

    const onSubmitform = async e => {
        e.preventDefault();
        console.log(Employee)
        await axios.post("http://localhost:9000/api/employees",Employee)
        history.push("/");

    }
    return (
        <div className="container">
            <br></br>
            <h1>Employee</h1>
            <form onSubmit={e => onSubmitform(e)}>
                <div className="form-group">
                    <label>Name</label>
                    <input type="text" className="form-control" onChange={e => onInputChange(e)} value={name} autoComplete="off" name="name" id="name" placeholder="Enter your name" />
                </div>
                <div className="form-group">
                    <label>City</label>
                    <input type="text" className="form-control" onChange={e => onInputChange(e)} value={city} autoComplete="off" name="city" id="city" placeholder="Enter your city" />
                </div>
                <div className="form-group">
                    <label>Mobile</label> 
                    <input type="text" className="form-control" onChange={e => onInputChange(e)} value={mobile} autoComplete="off" id="mobile" name="mobile" placeholder="Enter your mobile number" />
                </div>
                <div className="form-group">
                    <label>Email</label>
                    <input type="email" className="form-control" onChange={e => onInputChange(e)} value={email} autoComplete="off" id="email" name="email" placeholder="Enter your email address" />
                </div>
           
                <button type="submit" className="btn btn-primary">Submit</button>
            </form>
        </div>
    )
}

export default AddEmployee