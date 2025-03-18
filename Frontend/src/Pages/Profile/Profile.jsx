import React, { useEffect, useState } from "react";
import "./Profile.css";
import axios from "axios";
import Navbar from "../../Components/Navbar/Navbar";

function Profile() {
	const [user, setUser] = useState({
		name: "",
		email: "",
		role: "",
		contact_number: "",
		libid: "",
	});

	useEffect(() => {
		let isMounted = true; // Add a flag to check if the component is mounted
		const id = localStorage.getItem("user_id");
		if (id) {
			axios
				.get(`http://localhost:8080/user/${id}`)
				.then((response) => {
					if (isMounted) {
						const userData = response.data.data.User;
						// console.log(userData);
						setUser({
							name: userData.name,
							email: userData.email,
							role: userData.role,
							contact_number: userData.contact_number,
							libid: userData.libid,
						});
						// console.log("user data retreived.");
					}
				})
				.catch((error) => {
					if (isMounted) {
						console.error(
							"There was an error fetching the user data!",
							error
						);
					}
				});
		} else {
			console.error("No user ID found in localStorage");
		}
		return () => {
			isMounted = false; // Cleanup function to set the flag to false
		};
	}, []);
	return (
		<>
			<Navbar />
			<div className="profile-wrapper">
				<div className="profile-container">
					<h1>Profile</h1>
					<p>
						<strong>Name:</strong> {user.name}
					</p>
					<p>
						<strong>Email:</strong> {user.email}
					</p>
					<p>
						<strong>Contact Number:</strong> {user.contact_number}
					</p>
					<p>
						<strong>Role:</strong> {user.role}
					</p>
					{user.libid === "" && (
						<p>
							<strong>Lib Id:</strong> {user.libid}
						</p>
					)}
				</div>
			</div>
		</>
	);
}

export default Profile;
