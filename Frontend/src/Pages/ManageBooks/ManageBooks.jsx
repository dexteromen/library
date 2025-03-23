import React, { useState, useEffect } from "react";
import Navbar from "./../../Components/Navbar/Navbar";
import "./ManageBooks.css";
import { getRequests, returnBook } from "../../API/API";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

export default function ManageBooks() {
    const [AllRequests, setAllRequests] = useState([]);

    useEffect(() => {
        const fetchRequests = async () => {
            try {
                const res = await getRequests();
                const bookRequestDetails = res.data.data;
                // console.log("Book Request: ", bookRequestDetails);
                setAllRequests(bookRequestDetails);
            } catch (error) {
                console.log(error);
            }
        };
        fetchRequests();

        //Return Book also create
    }, []);

    // console.log("All-Requests::", AllRequests);
    const id = parseInt(localStorage.getItem("user_id"));

    // Replace with actual reader_id
    // const readerId = id;
    const readerId = id;

    const filteredRequests = AllRequests.filter(
        (request) => request.reader_id === readerId
    );

    console.log("Filtered Requests::", filteredRequests);

    // const handleReturnBook = async (requestId) => {
    //     try {
    //         const res = await returnBook(requestId);
    //         console.log("Return Book: ", res);
    //         // Reload the page
    //         setTimeout(() => {
    //             window.location.reload();
    //         }, 500);
    //     } catch (error) {
    //         console.log(error);
    //     }
    // };

    const handleReturnBook = async (isbn) => {
        try {
            const res = await returnBook(isbn);

            if (res) {
                // Process the data
                // var message = res.data.message;
                console.log(res);
                // toast.success(message);

                // Reload the page
                setTimeout(() => {
                    window.location.reload();
                }, 2500);
            } else {
                console.error("Response or data is undefined");
            }
        } catch (error) {
            console.error("Error fetching data", error);

            // if (error.response.statusText) {
            //     var errMessage = error.response.data.error;
            //     toast.error(errMessage);
            // }
        }
    };

    return (
        <>
            <Navbar />
            <div className="home-wrapper">
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
                <h1 className="manage-book-title">Manage Books</h1>
                <div className="book-wrapper">
                    {filteredRequests.length > 0 ? (
                        filteredRequests.map((book, index) => (
                            <div key={index} className="book-list">
                                <div className="book">
                                    <div className="book-label">
                                        ISBN: {""}{" "}
                                        <span className="book-value">
                                            {book.isbn}
                                        </span>
                                    </div>
                                    <div className="book-label">
                                        Issue Status: {""}{" "}
                                        <span className="book-value">
                                            {book.issue_status}
                                        </span>
                                    </div>
                                    <div className="book-label">
                                        Request Type: {""}{" "}
                                        <span className="book-value">
                                            {book.request_type}
                                        </span>
                                    </div>
                                    <div className="book-actions">
                                        {/* <button>Return Book</button> */}
                                        <button
                                            className="button"
                                            type="submit"
                                            onClick={() => {
                                                handleReturnBook(book.isbn);
                                                console.log(book.isbn);
                                                console.log(
                                                    "Return Book clicked"
                                                );
                                            }}
                                        >
                                            Return
                                            <svg
                                                viewBox="0 0 16 19"
                                                xmlns="http://www.w3.org/2000/svg"
                                            >
                                                <path d="M7 18C7 18.5523 7.44772 19 8 19C8.55228 19 9 18.5523 9 18H7ZM8.70711 0.292893C8.31658 -0.0976311 7.68342 -0.0976311 7.29289 0.292893L0.928932 6.65685C0.538408 7.04738 0.538408 7.68054 0.928932 8.07107C1.31946 8.46159 1.95262 8.46159 2.34315 8.07107L8 2.41421L13.6569 8.07107C14.0474 8.46159 14.6805 8.46159 15.0711 8.07107C15.4616 7.68054 15.4616 7.04738 15.0711 6.65685L8.70711 0.292893ZM9 18L9 1H7L7 18H9Z"></path>
                                            </svg>
                                        </button>
                                    </div>
                                </div>
                            </div>
                        ))
                    ) : (
                        <p>No books to display</p>
                    )}
                </div>
            </div>
        </>
    );
}
