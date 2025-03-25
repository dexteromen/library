import React, { useEffect, useState } from "react";
import "./Dashboard.css";
import Navbar from "../../Components/Navbar/Navbar";
import { useNavigate, Link } from "react-router-dom";
import { MdOutlineKeyboardBackspace } from "react-icons/md";
import {
	getUsers,
	getLibraries,
	getRequests,
	getIssues,
	approveAndIssueRequest,
	getUserById,
	deleteUserById,
} from "../../API/API";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

function AdminDashboard() {
	const navigate = useNavigate();
	const [admin, setAdmin] = useState([]);
	const [users, setUsers] = useState([]);
	const [libraries, setLibraries] = useState([]);
	const [AllRequests, setAllRequests] = useState([]);
	const [AllIssue, setAllIssue] = useState([]);
	const [update, setUpdate] = useState(false); // New state to trigger re-render

	useEffect(() => {
		const fetchUsers = async () => {
			try {
				const res = await getUsers();
				const userDetails = res.data.data.User;
				console.log("User Details: ", userDetails);
				setUsers(userDetails);
			} catch (error) {
				console.log("Error Fetching Users:", error);
			}
		};
		fetchUsers();
		const fetchAdmin = async () => {
			try {
				const user_id = localStorage.getItem("user_id");
				const parse_user_id = parseInt(user_id);
				const res = await getUserById(parse_user_id);
				const adminDetails = res.data.data.User;
				console.log("Admin Details: ", adminDetails);
				setAdmin(adminDetails);
			} catch (error) {
				console.log("Error Fetching Admin:", error);
			}
		};
		fetchAdmin();
		const fetchLibrary = async () => {
			try {
				const res = await getLibraries();
				const libraryDetails = res.data.data;
				console.log("Libraries: ", libraryDetails);
				setLibraries(libraryDetails);
			} catch (error) {
				console.log(error);
			}
		};
		fetchLibrary();
		const fetchRequests = async () => {
			try {
				const res = await getRequests();
				const bookRequestDetails = res.data.data;
				console.log("Book Request: ", bookRequestDetails);
				setAllRequests(bookRequestDetails);
			} catch (error) {
				console.log(error);
			}
		};
		fetchRequests();
		const fetchIssues = async () => {
			try {
				const res = await getIssues();
				const issuesDetails = res.data.data;
				console.log("All Issues: ", issuesDetails);
				setAllIssue(issuesDetails);
			} catch (error) {
				console.log(error);
			}
		};
		fetchIssues();
	}, [update]);

	const handleApproveRequest = async (req_Id) => {
		try {
			const res = await approveAndIssueRequest(req_Id);

			if (res) {
				// Process the data
				var message = res.data.message;
				console.log(message);
				setUpdate(!update);
				toast.success(message);

				// // Reload the page
				// setTimeout(() => {
				// 	window.location.reload();
				// }, 2500);

				// Trigger re-render by updating 'update' state
			} else {
				console.error("Response or data is undefined");
			}
		} catch (error) {
			console.error("Error fetching data", error);

			if (error.response.statusText) {
				var errMessage = error.response.data.error;
				toast.error(errMessage);
			}
		}
	};
	const handleDeleteUser = async (userId) => {
		try {
			const res = await deleteUserById(userId);

			if (res) {
				// Process the data
				var message = res.data.message;
				// console.log(res);
				setUpdate(!update);
				toast.success(message);

				// // Reload the page
				// setTimeout(() => {
				// 	window.location.reload();
				// }, 2500);

				// Trigger re-render by updating 'update' state
			} else {
				console.error("Response or data is undefined");
			}
		} catch (error) {
			console.error("Error in Deleting User", error);
			if (error.response.statusText) {
				var errMessage = error.response.data.error;
				toast.error(errMessage);
			}
		}
	};
	// const librarymap = new Map();
	// users.filter((user) => {
	// 	librarymap.set(user.name, user.lib_id);
	// });

	// const lib_owner = new Map();
	// libraries.filter((library) => {
	// 	librarymap.forEach((libid, ownerName) => {
	// 		// console.log("keys: ", ownerName + ", values: ", libid);
	// 		if (library.id == libid) {
	// 			lib_owner.set(library.name, ownerName);
	// 		}
	// 	});
	// });

	// console.log("Lib-Owner", lib_owner);

	return (
		<>
			<Navbar />
			<div className="dashboard-wrapper">
				<div className="admin-dashboard">
					<ToastContainer
						position="top-center"
						autoClose={2000}
						hideProgressBar={false}
						newestOnTop={false}
						closeOnClick
						rtl={false}
						pauseOnFocusLoss
						draggable
						pauseOnHover
					/>
					<h1 className="admin-dashboard__title">Admin Dashboard</h1>

					<div className="widget-container">
						<div className="widget">
							<h2 className="widget__title">Users</h2>
							<ul className="users__list">
								{users.map((user) => (
									<li key={user.id} className="users__item">
										{user.id} {user.name} ({user.email})
										{user.role !== admin.role && (
											<button
												className="button-delete"
												onClick={() =>
													handleDeleteUser(user.id)
												}
											>
												Delete User
											</button>
										)}
									</li>
								))}
							</ul>
						</div>

						<div className="widget">
							<h2 className="widget__title">Libraries</h2>
							{libraries.map((library) => (
								<div key={library.id} className="library">
									<h4 className="library__name">
										{library.name}
									</h4>
								</div>
							))}
						</div>

						<div className="widget">
							<h2 className="widget__title">Book Requests</h2>
							<div className="requests-container">
								{AllRequests.map((request) => (
									<div
										className="book-request"
										key={request.req_id}
									>
										<table>
											<tbody>
												<tr>
													<td className="book-table-label">
														ISBN:
													</td>
													<td className="book-table-value">
														{request.isbn}
													</td>
												</tr>
												<tr>
													<td className="book-table-label">
														Request Type:
													</td>
													<td className="book-table-value">
														{request.request_type}
													</td>
												</tr>
												<tr>
													<td className="book-table-label">
														Issue Status:
													</td>
													<td className="book-table-value">
														{request.issue_status ===
														"Approved And Issued"
															? "Issued"
															: "Pending"}
													</td>
												</tr>
												<tr>
													<td className="book-table-label">
														Reader ID:
													</td>
													<td className="book-table-value">
														{request.reader_id}
													</td>
												</tr>
												{request.approval_date && (
													<tr>
														<td className="book-table-label">
															Approval Date:
														</td>
														<td className="book-table-value">
															{request.approval_date.substring(
																0,
																10
															)}
														</td>
													</tr>
												)}
											</tbody>
										</table>
										{request.issue_status == "Pending" && (
											<button
												className="book-approve-btn"
												onClick={() =>
													handleApproveRequest(
														request.req_id
													)
												}
											>
												Approve and Issue
											</button>
										)}
									</div>
								))}
							</div>
						</div>

						<div className="widget">
							<h2 className="widget__title">Issued Books</h2>
							<div className="requests-container">
								{AllIssue.map((issue) => (
									<div
										className="book-request"
										key={issue.issue_id}
									>
										<table>
											<tbody>
												<tr>
													<td className="book-table-label">
														ISBN:
													</td>
													<td className="book-table-value">
														{issue.isbn}
													</td>
												</tr>
												<tr>
													<td className="book-table-label">
														Reader Id:
													</td>
													<td className="book-table-value">
														{issue.reader_id}
													</td>
												</tr>
												<tr>
													<td className="book-table-label">
														Issue Status:
													</td>
													<td className="book-table-value">
														{issue.issue_status}
													</td>
												</tr>
												<tr>
													<td className="book-table-label">
														Issue Date:
													</td>
													<td className="book-table-value">
														{issue.issue_date.substring(
															0,
															10
														)}
													</td>
												</tr>
												<tr>
													<td className="book-table-label">
														Expected Return Date:
													</td>
													<td className="book-table-value">
														{
															issue.expected_return_date
														}
													</td>
												</tr>
												{issue.return_date && (
													<tr>
														<td className="book-table-label">
															Return Date:
														</td>
														<td className="book-table-value">
															{issue.return_date}
														</td>
													</tr>
												)}
											</tbody>
										</table>
									</div>
								))}
							</div>
						</div>
					</div>
				</div>
				<div className="child-back-btn">
					<Link to="/home">
						<MdOutlineKeyboardBackspace size={45} />
					</Link>
				</div>
			</div>
		</>
	);
}

export default AdminDashboard;
