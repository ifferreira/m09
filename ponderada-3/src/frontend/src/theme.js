import { createTheme } from "@mui/material/styles";

const theme = createTheme({
  palette: {
    primary: {
      main: "#00bcd4",  // Cor primária vibrante
    },
    secondary: {
      main: "#ff4081",  // Cor secundária vibrante
    },
    background: {
      default: "#fafafa",  // Cor de fundo clara
      paper: "#ffffff",    // Cor de fundo dos papéis
    },
    text: {
      primary: "#333333",  // Texto principal mais escuro
      secondary: "#555555",  // Texto secundário mais suave
    },
  },
  typography: {
    fontFamily: [
      '"Roboto Slab"',
      "serif",  // Alterando a fonte para dar um toque mais sofisticado
    ].join(","),
    h1: {
      fontWeight: 700,
    },
    h2: {
      fontWeight: 600,
    },
    body1: {
      fontWeight: 400,
      lineHeight: 1.6,
    },
  },
  components: {
    MuiButton: {
      styleOverrides: {
        root: {
          textTransform: "none",
          borderRadius: 12,  // Borda mais arredondada
          padding: "8px 16px",  // Aumentando o padding para botões mais espaçados
          backgroundColor: "#00bcd4", // Cor personalizada para o botão
          color: "#fff", // Cor do texto nos botões
          "&:hover": {
            backgroundColor: "#008c9e", // Hover com cor mais escura
          },
        },
      },
    },
    MuiCard: {
      styleOverrides: {
        root: {
          borderRadius: 16,  // Borda mais arredondada para cards
          boxShadow: "0 6px 18px rgba(0,0,0,0.15)",  // Sombra mais intensa para cards
          transition: "all 0.3s ease",  // Transição suave ao passar o mouse
          "&:hover": {
            transform: "scale(1.05)",  // Aumenta o card ao passar o mouse
            boxShadow: "0 12px 36px rgba(0,0,0,0.2)",  // Sombra mais forte no hover
          },
        },
      },
    },
  },
});

export default theme;
