import React from "react";
import { NavLink, Link, useNavigate } from "react-router-dom";
import "./Navbar.css";
import logo from "/library.svg"; //in public folder

export default function Navbar() {
	const navigate = useNavigate();
	// const isLoggedIn = true;
	const isLoggedIn = false;
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
					{isLoggedIn ? (
						<>
							<img
								src="https://avatar.iran.liara.run/public/boy"
								alt="profile-image"
							/>
							<button
								// className="btn-login"
								className="button-38"
								onClick={() => navigate("/temp")}
							>
								LOGOUT
							</button>
						</>
					) : (
						<>
							<button
								// className="btn-login"
								className="button-38"
								onClick={() => navigate("/login")}
							>
								LOGIN
							</button>
							<button
								className="button-38"
								onClick={() => navigate("/signup")}
							>
								SIGNUP
							</button>
						</>
					)}
				</div>
			</div>
		</>
	);
}
