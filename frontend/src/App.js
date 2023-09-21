import "./App.css";
import { useCookies } from "react-cookie";
import { useEffect, useState } from "react";
import {
  Input,
  Form,
  FormGroup,
  Label,
  Container,
  Row,
  Button,
  Accordion,
  AccordionItem,
  AccordionHeader,
  AccordionBody,
  Spinner,
} from "reactstrap";
import Navigation from "./Navigation";

const login = `${process.env.REACT_APP_BACKEND}/api/auth/google/login`;
const googleApis = `https://www.googleapis.com/oauth2/v2/userinfo`;

function App() {
  const [cookies, removeCookies] = useCookies(["token"]);
  const [loading, setLoading] = useState(false);
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);
  const [authenticated, setAuthenticated] = useState(false);
  const [profile, setProfile] = useState(null);
  const [open, setOpen] = useState(false);
  const [validation, setValidation] = useState(null);

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

  const toggle = (id) => {
    if (open === id) {
      setOpen();
    } else {
      setOpen(id);
    }
  };

  const onClick = () => {
    removeCookies("token");
    window.location.reload();
  };

  if (authenticated) {
    return (
      <div>
        <Navigation author={`Hello ${profile?.email}`}>
          <Button onClick={onClick} color="danger">
            logout
          </Button>
        </Navigation>
        <br />
        <Container>
          <Row>
            <Form
              onSubmit={(e) => {
                let objective = e.target.objective.value || "";
                let translate = e.target.translate.value || "english";

                setLoading(true);

                if (objective === "") {
                  setValidation("Please fill objective");
                } else {
                  fetch(
                    `${process.env.REACT_APP_BACKEND}/api/v1/okr-generator`,
                    {
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
                    },
                  )
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
                }
              }}
            >
              <FormGroup>
                <Label>Objective</Label>
                <Input type="text" name="objective" placeholder="objective" />
              </FormGroup>
              <FormGroup>
                <Label>Translate</Label>
                <Input type="text" name="translate" placeholder="translate" />
              </FormGroup>
              <Input type="submit" disabled={loading} />
            </Form>
          </Row>
        </Container>
        <Row>
          {loading && (
            <p style={{ "text-align": "center", padding: "10px" }}>
              Loading...
            </p>
          )}
          {error && (
            <p style={{ "text-align": "center", padding: "10px" }}>{error}</p>
          )}
          {validation && (
            <p style={{ "text-align": "center", padding: "10px" }}>{error}</p>
          )}
        </Row>
        <Container>
          <Row>
            {(function() {
              if (data) {
                return (
                  <Accordion flush open={open} toggle={toggle}>
                    <h2>{data?.objective}</h2>
                    {data?.key_results.map((kr, index) => (
                      <AccordionItem key={String(index)}>
                        <AccordionHeader targetId={`${index}`}>
                          {kr?.key_result}
                        </AccordionHeader>
                        <AccordionBody accordionId={`${index}`}>
                          <small>{kr?.key_result}</small>
                        </AccordionBody>
                      </AccordionItem>
                    ))}
                  </Accordion>
                );
              }
            })()}
          </Row>
        </Container>
      </div>
    );
  } else {
    return (
      <div>
        <br />
        <br />
        <Container>
          <Row md={{ offset: 3, size: 6 }} sm="12">
            <Button color="primary" href={login}>
              Google Auth
            </Button>
          </Row>
        </Container>
      </div>
    );
  }
}

export default App;
