import React from "react";
import { NavLink, Link, useNavigate } from "react-router-dom";
import "./Navbar.css";
import logo from "/public/library.svg";

export default function Navbar() {
	const navigate = useNavigate();
	return (
		<>
			<div className="nav-bar">
				<div className="logo-wrapper">
					<div className="image">
						<Link to="/">
							<img src={logo} alt="Logo" />
						</Link>
					</div>
					<div className="logo-name">Library Management System</div>
				</div>
				<div className="links">
					<NavLink to="/" activeClassName="active">
						Home
					</NavLink>
					<Link to="/">Libraries</Link>
					<Link to="/">About</Link>
				</div>
				<div className="profile">
					profile-image
					<button
						className="btn-login"
						onClick={() => navigate("/signin")}
					>
						Signin
					</button>
					<button
						className="btn-signup"
						onClick={() => navigate("/signup")}
					>
						Signup
					</button>
				</div>
			</div>
		</>
	);
}
