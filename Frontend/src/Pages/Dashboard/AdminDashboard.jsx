import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import "./Dashboard.css";
import Navbar from "../../Components/Navbar/Navbar";
import { getUsers, getLibraries, getRequests, getIssues } from "../../API/API";
import { Link } from "react-router-dom";
import { MdOutlineKeyboardBackspace } from "react-icons/md";

function AdminDashboard() {
	const navigate = useNavigate();
	const [libraries, setLibraries] = useState([]);
	const [bookRequests, setBookRequests] = useState([]);
	const [issuedBooks, setIssuedBooks] = useState([]);
	const [users, setUsers] = useState([]);

	useEffect(() => {
		// Dummy data for libraries
		const dummyLibraries = [
			{
				id: 1,
				name: "Central Library",
				owner: "John Doe",
				readersCount: 120,
				books: [
					{ id: 1, title: "Book One" },
					{ id: 2, title: "Book Two" },
				],
			},
			{
				id: 2,
				name: "Community Library",
				owner: "Jane Smith",
				readersCount: 80,
				books: [
					{ id: 3, title: "Book Three" },
					{ id: 4, title: "Book Four" },
				],
			},
		];
		// setLibraries(dummyLibraries);

		// Dummy data for book requests
		const dummyBookRequests = [
			{ id: 1, bookTitle: "Book One", readerName: "Alice" },
			{ id: 2, bookTitle: "Book Three", readerName: "Bob" },
		];
		setBookRequests(dummyBookRequests);

		// Dummy data for issued books
		const dummyIssuedBooks = [
			{
				id: 1,
				bookTitle: "Book Two",
				readerName: "Charlie",
				issueDate: "2025-03-01",
			},
			{
				id: 2,
				bookTitle: "Book Four",
				readerName: "David",
				issueDate: "2025-03-05",
			},
		];
		setIssuedBooks(dummyIssuedBooks);

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
				// setBookRequests(bookRequestDetails);
			} catch (error) {
				console.log(error);
			}
		};
		fetchRequests();
		const fetchIssues = async () => {
			try {
				const res = await getIssues();
				const issuesDetails = res.data.data;
				console.log("Book Issued: ", issuesDetails);
				// setIssuedBooks(issuesDetails);
			} catch (error) {
				console.log(error);
			}
		};
		fetchIssues();
	}, []);

	const handleApproveRequest = (requestId) => {
		// Approve and issue the book request
		const approvedRequest = bookRequests.find(
			(request) => request.id === requestId
		);
		if (approvedRequest) {
			const newIssuedBook = {
				id: issuedBooks.length + 1,
				bookTitle: approvedRequest.bookTitle,
				readerName: approvedRequest.readerName,
				issueDate: new Date().toISOString().split("T")[0],
			};
			setIssuedBooks([...issuedBooks, newIssuedBook]);
			setBookRequests(
				bookRequests.filter((request) => request.id !== requestId)
			);
		}
	};

	return (
		<>
			<Navbar />
			<div className="dashboard-wrapper">
				<div className="admin-dashboard">
					<h1 className="admin-dashboard__title">Admin Dashboard</h1>

					<div className="widget-container">
						<div className="widget">
							<h2 className="widget__title">Book Requests</h2>
							<ul className="book-requests__list">
								{bookRequests.map((request) => (
									<li
										key={request.id}
										className="book-requests__item"
									>
										{request.bookTitle} requested by{" "}
										{request.readerName}
										<button
											className="book-requests__approve-button"
											onClick={() =>
												handleApproveRequest(request.id)
											}
										>
											Approve and Issue
										</button>
									</li>
								))}
							</ul>
						</div>

						<div className="widget">
							<h2 className="widget__title">Issued Books</h2>
							<ul className="issued-books__list">
								{issuedBooks.map((issue) => (
									<li
										key={issue.id}
										className="issued-books__item"
									>
										{issue.bookTitle} issued to{" "}
										{issue.readerName} on {issue.issueDate}
									</li>
								))}
							</ul>
						</div>

						<div className="widget">
							<h2 className="widget__title">Users</h2>
							<ul className="users__list">
								{users.map((user) => (
									<li key={user.id} className="users__item">
										{user.name} ({user.email})
										<button>Delete User</button>
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
