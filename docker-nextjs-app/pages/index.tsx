import { useEffect, useState } from 'react'
import axios from 'axios'


//////////////////////* uncomment below code to use api that gives an hello message to next js frontend  */////////////////////////////////////////////////////

// const Home = () => {
//   const [message, setMessage] = useState('')
//   const [error, setError] = useState<string>('')

//   useEffect(() => {
//     const fetchMessage = async () => {
//       try {
//         const response = await axios.get('http://localhost:1190/')
//         setMessage(response.data.message)
//       } catch (error: any) { // Specify the type of error as 'any' or 'Error'
//         setError(error.message)
//         console.error('Error fetching the message:', error)
//       }
//     }

//     fetchMessage()
//   }, [])

//   return (
//     <div>
//       {error ? (
//         <h1>Error: {error}</h1>
//       ) : (
//         <h1>Response from Go server: {message}</h1>
//       )}
//     </div>
//   );
// };

// export default Home
// pages/index.tsx


//////////////////////* uncomment below code to use api that reads wordpress post info using wordpress api's & gives reponse to the nextjs frontend */////////////////////////////////////////////////////
interface Post {
  id: number;
  title: {
    rendered: string;
  };
  content: {
    rendered: string;
  };
}

const Home = () => {
  const [posts, setPosts] = useState<Post[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const response = await axios.get('http://localhost:1190/api/posts');
        setPosts(response.data);
      } catch (error: any) {
        setError(error.message);
      }
    };

    fetchPosts();
  }, []);

  return (
    <div>
      {error && <h1>Error: {error}</h1>}
      <ul>
        {posts.map(post => (
          <li key={post.id}>
            <h2>{post.title.rendered}</h2>
            <div dangerouslySetInnerHTML={{ __html: post.content.rendered }} />
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Home;
