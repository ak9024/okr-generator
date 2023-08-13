import "./App.css";
import { useCookies } from "react-cookie";
import { useEffect, useState } from "react";

const login = `${process.env.REACT_APP_BACKEND}/api/auth/google/login`;
const googleApis = `https://www.googleapis.com/oauth2/v2/userinfo`;

function App() {
  const [cookies, removeCookies] = useCookies(["token"]);
  const [loading, setLoading] = useState(false);
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);
  const [authenticated, setAuthenticated] = useState(false);
  const [profile, setProfile] = useState(null);

  useEffect(() => {
    fetch(`${googleApis}?access_token=${cookies?.token}`)
      .then((res) => res.json())
      .then((res) => {
        if (profile?.error?.status !== "UNAUTHENTICATED") {
          setProfile(res);
          setAuthenticated(true);
        } else {
          setAuthenticated(false);
        }
      })
      .catch((err) => setError(err));
  }, [authenticated, cookies?.token]);

  const onClick = () => {
    removeCookies("token");
    window.location.reload();
  };

  if (authenticated) {
    return (
      <div>
        <p>Hello {profile?.email}</p>
        <form
          onSubmit={(e) => {
            let objective = e.target.objective.value;
            let translate = e.target.translate.value;

            setLoading(true);

            fetch(`${process.env.REACT_APP_BACKEND}/api/v1/okr-generator`, {
              method: "POST",
              mode: "cors",
              headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${cookies?.token}`,
              },
              body: JSON.stringify({
                objective,
                translate,
              }),
            })
              .then((res) => res.json())
              .then((res) => {
                setData(res);
                setLoading(false);
              })
              .catch((err) => {
                setError(err);
                setLoading(false);
              });

            e.preventDefault();
            e.currentTarget.reset();
          }}
        >
          <input type="text" name="objective" placeholder="objective" />
          <input type="text" name="translate" placeholder="translate" />
          <input type="submit" disabled={loading} />
        </form>
        {loading && <p>Loading</p>}
        {error && <p>{error}</p>}
        <div>
          {(function () {
            if (data) {
              return (
                <div>
                  <h2>{data?.objective}</h2>
                  {data?.key_results.map((kr, index) => (
                    <div key={String(index)}>
                      <p>{kr?.key_result}</p>
                    </div>
                  ))}
                </div>
              );
            }
          })()}
        </div>
        <button onClick={onClick}>logout</button>
      </div>
    );
  } else {
    return (
      <div>
        <a href={login}>Login</a>
      </div>
    );
  }
}

export default App;
