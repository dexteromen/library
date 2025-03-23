// import React from "react";
// import Navbar from "../../Components/Navbar/Navbar";
// // import "./Temp.css";

// function Temp() {
// 	return (
// 		<>
// 			<Navbar />
// 			<div>Temp</div>
// 		</>
// 	);
// }

// export default Temp;

import React, { useEffect, useState } from "react";
import { getUsers } from "../../API/API";

const Temp = () => {
	const [users, setUsers] = useState([]);

	useEffect(() => {
		const fetchUsers = async () => {
			try {
				const response = await getUsers();
				const userD = response.data.data.User;
				// console.log(userD);
				setUsers(userD);
			} catch (error) {
				console.error("Error fetching users:", error);
			}
		};

		fetchUsers();
	}, []);

	const handleSignUp = async (userData) => {
		try {
			const response = await signUp(userData);
			console.log("User signed up:", response.data);
		} catch (error) {
			console.error("Error signing up:", error);
		}
	};

	return (
		<div>
			<h1>Users</h1>
			<ul>
				{users.map((user) => (
					<li key={user.id}>
						{user.name} {user.email}
					</li>
				))}
			</ul>
			{/* Add form or button to handle sign up */}
		</div>
	);
};

export default Temp;
