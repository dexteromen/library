import React from "react";
import { NavLink, Link, useNavigate } from "react-router-dom";
import "./Navbar.css";
import logo from "/library.svg"; //in public folder

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
					<div className="logo-name">XenonLibrary</div>
				</div>
				<div className="links">
					<NavLink to="/">Home</NavLink>
					<Link to="/">Libraries</Link>
					<Link to="/">About</Link>
				</div>
				<div className="profile">
					<img
						src="https://avatar.iran.liara.run/public/boy"
						alt="profile-image"
					/>
					<button
						// className="btn-login"
						className="button-38"
						onClick={() => navigate("/signin")}
					>
						Signin
					</button>
					<button
						className="button-38"
						onClick={() => navigate("/signup")}
					>
						Signup
					</button>
				</div>
			</div>
		</>
	);
}
