import React, { useEffect, useState } from "react";
import { NavLink, Link, useNavigate } from "react-router-dom";
import "./Navbar.css";
import logo from "/library.svg"; //in public folder

export default function Navbar() {
	const navigate = useNavigate();
	const [isLoggedIn, setIsLoggedIn] = useState(false);

	useEffect(() => {
		const token = localStorage.getItem("token");
		if (token) {
			setIsLoggedIn(true);
		}
	}, []);

	const handleLogout = () => {
		localStorage.removeItem("token");
		localStorage.removeItem("user_id");
		setIsLoggedIn(false);
		navigate("/login");
	};

	return (
		<>
			<div className="nav-bar">
				<div className="logo-wrapper">
					<div className="image">
						<Link to="/">
							<img src={logo} alt="Logo" />
						</Link>
					</div>
					<div className="logo-name">Z-Lib</div>
				</div>
				{/* <div className="links">
					<NavLink to="/">Home</NavLink>
					{isLoggedIn && (
						<Link to="/create-library">Create Library</Link>
					)}
					{isLoggedIn && <Link to="/create-book">Create Book</Link>}
					{isLoggedIn && <Link to="/dashboard">dashboard</Link>}
					{isLoggedIn && <Link to="/temp">Temp</Link>}
				</div> */}
				<div className="profile">
					{isLoggedIn ? (
						<>
							<Link to="/profile">
								<img
									src="https://avatar.iran.liara.run/public/boy"
									alt="profile-image"
								/>
							</Link>
							<button
								className="button-38"
								onClick={handleLogout}
							>
								LOGOUT
							</button>
						</>
					) : (
						<>
							<button
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
