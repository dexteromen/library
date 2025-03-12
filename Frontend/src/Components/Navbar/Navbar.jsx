import React from "react";
import "./Navbar.css";

export default function Navbar() {
	return (
		<>
			<div className="nav-bar">
				<div className="logo">logo</div>
				<div className="links">
					<a href="#">Home</a>
					<a href="#">Libraries</a>
					<a href="#">About</a>
				</div>
				<div className="profile">
					profile
					<button className="btn-login">Login</button>
					<button className="btn-signup">Signup</button>
				</div>
			</div>
		</>
	);
}
