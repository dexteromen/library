import axios from 'axios';

const API_URL = 'http://localhost:8080';

// Auth Routes
export const signUp = async (userData) => {
    return await axios.post(`${API_URL}/signup`, userData);
};

export const signIn = async (credentials) => {
    return await axios.post(`${API_URL}/signin`, credentials);
};

export const signOut = async () => {
    var token = localStorage.getItem("token");
    return await axios.post(`${API_URL}/signout`,{
        headers: {
            'Authorization':`Bearer ${token}`
        }
    });
};

// User Routes
export const getUsers = async () => {
    return await axios.get(`${API_URL}/users`);
};

export const getUserById = async (id) => {
    return await axios.get(`${API_URL}/user/${id}`);
};

export const deleteUserById = async (id) => {
    var token = localStorage.getItem("token");
    return await axios.delete(`${API_URL}/user/${id}`,{
        headers: {
            'Authorization':`Bearer ${token}`
        }
    });
};

// Library Routes
export const getLibraries = async () => {
    return await axios.get(`${API_URL}/library`);
};

export const createLibrary = async (lib_name) => {
    var token = localStorage.getItem("token");
    return await axios.post(`${API_URL}/library`, {
        "name": lib_name,
    },{
        headers: {
            'Authorization':`Bearer ${token}`
        }
    });
};

// Book Routes
export const searchBooks = async (query) => {
    return await axios.get(`${API_URL}/search`, { params: { q: query } });
};

export const getBooks = async () => {
    return await axios.get(`${API_URL}/books`);
};

export const getBookById = async (id) => {
    return await axios.get(`${API_URL}/book/${id}`);
};

export const createBook = async (bookData) => {
    var token = localStorage.getItem("token");
    // console.log(bookData)
    return await axios.post(`${API_URL}/book`, {
        "isbn": bookData.isbn,
        "title": bookData.title,
        "authors": bookData.authors,
        "publisher": bookData.publisher,
        "version": bookData.version,
        "total_copies": parseInt(bookData.total_copies,10),
        "available_copies": parseInt(bookData.available_copies,10),
    } ,{
        headers: {
            'Authorization':`Bearer ${token}`
        }
    });
};

export const updateBookById = async (id, bookData) => {
    var token = localStorage.getItem("token");
    return await axios.put(`${API_URL}/book/${id}`, bookData);
};

export const deleteBookById = async (id) => {
    var token = localStorage.getItem("token");
    return await axios.delete(`${API_URL}/book/${id}`);
};

// Request Issues Return Routes
export const getRequests = async () => {
    var token = localStorage.getItem("token");
    return await axios.get(`${API_URL}/requests`,{
        headers: {
            'Authorization':`Bearer ${token}`
        }
    });
};

export const getIssues = async () => {
    var token = localStorage.getItem("token");
    return await axios.get(`${API_URL}/issues`,{
        headers: {
            'Authorization':`Bearer ${token}`
        }
    });
};

export const createRequest = async (requestData) => {
    var token = localStorage.getItem("token");
    return await axios.post(`${API_URL}/request`, requestData,{
        headers: {
            'Authorization':`Bearer ${token}`
        }
    });
};

export const approveAndIssueRequest = async (id) => {
    var token = localStorage.getItem("token");
    return await axios.put(`${API_URL}/approve-issue/${id}`);
};

export const returnBook = async (id) => {
    var token = localStorage.getItem("token");
    return await axios.put(`${API_URL}/return/${id}`);
};

export const getProfile = async () => {
    var token = localStorage.getItem("token");
    return await axios.get(`${API_URL}/profile`,{
        headers: {
            'Authorization':`Bearer ${token}`
        }
    });
};

export const getProfileByToken = async () => {
    var token = localStorage.getItem("token");
    return await axios.get(`${API_URL}/profile-by-token`,{
        headers: {
            'Authorization':`Bearer ${token}`
        }
    });
};

// export const refreshToken = async () => {
//     var token = localStorage.getItem("token");
//     const response =  await axios.post(`${API_URL}/refresh-token`,{}, {
//         headers: {
//             'Authorization':`Bearer ${token}`
//         }
//     });

//     const newToken = response.data.token;
//     localStorage.setItem("token", newToken);

//     return response;
// };

export const refreshToken = async () => {
    try {
        var token = localStorage.getItem("token");
        const response = await axios.post(`${API_URL}/refresh-token`, {}, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });

        const newToken = response.data.data.token;
        console.log(newToken);
        localStorage.setItem("token", newToken);

        return response;
    } catch (error) {
        console.error("Error refreshing token:", error);
        throw error;
    }
};