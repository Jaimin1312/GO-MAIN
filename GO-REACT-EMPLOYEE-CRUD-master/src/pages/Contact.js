
import React from 'react'
function Contact() {
    return (
        <div className="container">
 
            <br />
            <h3 className="font-weight-bold">Query Form</h3>
            <br/>
            <form>
                <div class="form-group">
                    <label for="exampleInputEmail1">Email address</label>
                    <input type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Enter email" />
                    <small id="emailHelp" class="form-text text-muted">We'll never share your email with anyone else.</small>
                </div>
                <div class="form-group">
                    <label for="exampleInputPassword1">Password</label>
                    <input type="password" class="form-control" id="exampleInputPassword1" placeholder="Password" />
                </div>
                <div class="form-check">
                    <input type="checkbox" class="form-check-input" id="exampleCheck1" />
                    <label class="form-check-label" for="exampleCheck1">Check me out</label>
                </div>
                <button type="submit" class="btn btn-primary">Submit</button>
            </form>
            <br /> <br />
            <div className="py-4">
                <h3 className="font-weight-bold"> CONTACT ADDRESS</h3>
                <br /> <br />
                <div className="card">
                    <ul className="list-group list-group-flush">
                        <li className="list-group-item"> Dummy Address</li>
                        <li className="list-group-item"> Lorem Ipsum Sit Amet</li>
                        <li className="list-group-item">Dummy Pin</li>
                        <li className="list-group-item">Dummy place</li>
                        <li className="list-group-item">Telephone : 1-800-123-4567</li>
                        <li className="list-group-item"> Email : info@dummy.com</li>
                    </ul>
                </div>
                <br /> <br />
            </div>
        </div>
    );
}

export default Contact;
