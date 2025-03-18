import React, { useEffect, useState } from "react";
import { NavLink, Link, useNavigate } from "react-router-dom";
import "./Navbar.css";
import logo from "/library.svg"; //in public folder
import axios from "axios";

export default function Navbar() {
	const navigate = useNavigate();
	const [isLoggedIn, setIsLoggedIn] = useState(false);
	const [userRole, setUserRole] = useState("reader");

	useEffect(() => {
		const token = localStorage.getItem("token");
		const id = localStorage.getItem("user_id");

		//getting user role
		if (id && token) {
			axios
				.get(`http://localhost:8080/user/${id}`)
				.then((response) => {
					const userData = response.data.data.User;
					setUserRole(userData.role);
					setIsLoggedIn(true);
					localStorage.setItem("userRole", userRole);
					// console.log(userRole);
				})
				.catch((error) => {
					console.error(
						"There was an error fetching the user data!",
						error
					);
				});
		}
		// else {
		// 	console.error("No token found in localStorage");
		// 	console.error("No user ID found in localStorage");
		// }
	}, []);
	// [navigate]
	// );

	const handleLogout = () => {
		localStorage.removeItem("token");
		localStorage.removeItem("user_id");
		localStorage.removeItem("userRole");
		setUserRole("Logout");
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
				<div className="links">
					<NavLink to="/">Home</NavLink>
					{/* <Link to="/login">login</Link>
					<Link to="/signup">signup</Link>
					<Link to="/create-library">Create Library</Link>
					<Link to="/create-book"> Create Book</Link>
					<Link to="/dashboard">dashborad</Link>
					<Link to="/temp">Temp</Link> */}
					{isLoggedIn && (
						<Link to="/create-library">Create Library</Link>
					)}
					{isLoggedIn && <Link to="/create-book">Create Book</Link>}
					{isLoggedIn && <Link to="/dashboard">dashboard</Link>}
					{isLoggedIn && userRole === "admin" && (
						<Link to="/temp">Temp</Link>
					)}
				</div>
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
