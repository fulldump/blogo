package statics

var Files = map[string]string{
	"article.html": `e3tkZWZpbmUgImNvbnRlbnQifX0KCQkJPGgyPnt7LmFydGljbGUuVGl0bGV9fTwvaDI+CgkJCTxwPnt7LmFydGljbGUuQ29udGVudH19PC9wPgp7e2VuZH19Cgo=`,
	"favicon.ico": `AAABAAEAEBAAAAEAIABoBAAAFgAAACgAAAAQAAAAIAAAAAEAIAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAANnJZCDh0meF4dNmx+LTZ+/h0mb94dJn/+HSZ//h0mb94tNn7+HTZsfh0meF2clkIAAAAAAAAAAAAAAAAOHSZWvi0mf54dJn/sbRo/+t0dv/otHx/6LS9v+i0vb/otHx/63R2//G0aP/4dJn/uLSZ/nh0mVrAAAAAN/QZkXi02f/4tNn/8LRq/+i0vb/otL2/6LS9v+i0vb/otL2/6LS9v+i0vb/otL2/8LRq//i02f/4tNn/9/QZkXi0mfF4tNn/+LTZ/+r0d//otL2/6LS9v+i0vb/otL2/6LS9v+i0vb/otL2/6LS9v+r0d//4tNn/+LTZ//i0mfF4dNn9eLTZ//i02f/o9Hw/6LS9v+i0vb/otL2/6LS9v+i0vb/otL2/6LS9v+i0vb/o9Hw/+LTZ//i02f/4dNn9eHSZ/Li02f/4tNn/6rR4f+i0vb/otL2/6LS9v+i0vb/otL2/6LS9v+i0vb/otL2/6rR4f/i02f/4tNn/+HSZ/Lh0ma14tNn/+LTZ//B0q3/otL2/6LS9v+i0vb/otL2/6LS9v+i0vb/otL2/6LS9v/B0q3/4tNn/+LTZ//h0ma13NBlKuHSZ+3i02f/4NJo/8XRpf+s0N3/pNHv/6LS9v+i0vb/pNHv/6zQ3f/F0aX/4dNm++LTZ//h0mft3NBlKgAAAADY0Ggf4dJml+LSZ9fh02f14tNn/+LTZ//i02f/4tNn/+LTZ//i02f/4dNn9eLSZ9fh0maX2NBoHwAAAAAAAAAAAAAAAN3MZh3h0WaV2dJ4/8HRrv+r0d//otHx/6LR8f+r0d//wdGu/9nSeP/h0WaV3cxmHQAAAAAAAAAAAAAAAN/RZjbi0mbx4tNn/7vRvP+i0vb/otL2/6LS9v+i0vb/otL2/6LS9v+70bz/4tNn/+LSZvHf0WY2AAAAAAAAAADh02bT4tNn/+LTZ/+o0eX/otL2/6LS9v+i0vb/otL2/6LS9v+i0vb/qNHl/+LTZ//i02f/4dNm0wAAAAAAAAAA4tJm+OLTZ//i02f/pdHs/6LS9v//////otL2/6LS9v//////otL2/6XR7P/i02f/4tNn/+LSZvgAAAAAAAAAAOLTZszi02f/4tNn/7jRwv//////AAAA////////////AAAA//////+40cL/4tNn/+LTZ//i02bMAAAAAAAAAADezmc94dJm8+LTZ//e0m3/vNG4//////+i0fT/otH0//////+80bj/3tJt/+LTZ//h0mbz3s5nPQAAAAAAAAAAAAAAANjQaB/i02aI4dJmyeHSZ+3h02b74dJm/+HSZv/h02b74dJn7eHSZsni02aI2NBoHwAAAAAAAAAA4AcAAMADAACAAQAAAAAAAAAAAAAAAAAAAAAAAIABAADAAwAA4AcAAMADAACAAQAAgAEAAIABAADAAwAA4AcAAA==`,
	"index.html": `e3tkZWZpbmUgImNvbnRlbnQifX0KCgkJCXt7cmFuZ2UgLmFydGljbGVzfX0KCQkJPGRpdiBpZD0ie3suSWR9fSI+CgkJCQk8aDI+PGEgaHJlZj0iL2Eve3suSWR9fSI+e3suVGl0bGV9fTwvYT4gPGJ1dHRvbiBjbGFzcz0iYnV0dG9uLXJlbW92ZSIgb25jbGljaz0icmVtb3ZlQXJ0aWNsZSgne3suSWR9fScpIj5Cb3JyYXI8L2J1dHRvbj4gPC9oMj4KCQkJCTxwPnt7LkNvbnRlbnR9fTwvcD4KCQkJPC9kaXY+CgkJCXt7ZW5kfX0KCgkJCXt7aWYgLnVzZXJ9fQoJCQk8ZGl2IGNsYXNzPSJjcmVhdGUtZm9ybSI+CgkJCQk8aW5wdXQgdHlwZT0idGV4dCIgaWQ9ImNyZWF0ZS1mb3JtLXRpdGxlIiBwbGFjZWhvbGRlcj0iVMOtdHVsbyI+CgkJCQk8dGV4dGFyZWEgdHlwZT0idGV4dCIgaWQ9ImNyZWF0ZS1mb3JtLWNvbnRlbnQiIHBsYWNlaG9sZGVyPSJDb250ZW5pZG8uLi4iPjwvdGV4dGFyZWE+CgoJCQkJPGRpdiBzdHlsZT0idGV4dC1hbGlnbjogY2VudGVyOyI+CgkJCQkJPGJ1dHRvbiBpZD0iY3JlYXRlLWZvcm0tYnV0dG9uIiBjbGFzcz0iYnV0dG9uLWNyZWF0ZSIgb25jbGljaz0iY3JlYXRlQXJ0aWNsZSgpIj5DcmVhcjwvYnV0dG9uPgoJCQkJPC9kaXY+CgkJCTwvZGl2PgoJCQl7e2Vsc2V9fQoJCQk8ZGl2IHN0eWxlPSJ0ZXh0LWFsaWduOiBjZW50ZXI7Ij5JZGVudGlmw61jYXRlIHBhcmEgY3JlYXIgYXJ0w61jdWxvczwvZGl2PgoJCQl7e2VuZH19CgoJCTxkaXY+Cnt7ZW5kfX0KCg==`,
	"template.html": `PCFET0NUWVBFIGh0bWw+CjxodG1sPgoJPGhlYWQ+CgkJPG1ldGEgY2hhcnNldD0idXRmLTgiPgoJCTx0aXRsZT5CbG9HbzwvdGl0bGU+CgkJPG1ldGEgbmFtZT0idmlld3BvcnQiIGNvbnRlbnQ9IndpZHRoPWRldmljZS13aWR0aCwgaW5pdGlhbC1zY2FsZT0xIj4KCQk8bGluayByZWw9Imljb24iIHR5cGU9ImltYWdlL3gtaWNvbiIgY2xhc3M9ImpzLXNpdGUtZmF2aWNvbiIgaHJlZj0iL2Zhdmljb24uaWNvIj4KCQk8c3R5bGU+CgkJCWh0bWwgewoJCQkJZm9udC1zaXplOiAxMjAlOwoJCQl9CgoJCQkuY29udGVudCB7CgkJCQltYXgtd2lkdGg6IDgwMHB4OwoJCQkJbWFyZ2luOiBhdXRvOwoJCQl9CgoJCQloMSB7CgkJCQljb2xvcjogIzMwMzA2MDsKCQkJCXRleHQtYWxpZ246IGNlbnRlcjsKCQkJfQoKCQkJaDIgewoJCQkJY29sb3I6ICMzMDMwQTA7CgkJCX0KCgkJCS5idXR0b24tcmVtb3ZlIHsKCQkJCWJhY2tncm91bmQtY29sb3I6IHJlZDsKCQkJCWNvbG9yOndoaXRlOwoJCQkJYm9yZGVyOiBzb2xpZCAjNjYwMDAwIDFweDsKCQkJCWJvcmRlci1yYWRpdXM6IDNweDsKCQkJCWRpc3BsYXk6IGlubGluZS1ibG9jazsKCQkJCWN1cnNvcjogcG9pbnRlcjsKCQkJfQoKCQkJLmJ1dHRvbi1jcmVhdGUgewoJCQkJYmFja2dyb3VuZC1jb2xvcjogYmx1ZTsKCQkJCWNvbG9yOndoaXRlOwoJCQkJYm9yZGVyOiBzb2xpZCAjMDAwMDY2IDFweDsKCQkJCWJvcmRlci1yYWRpdXM6IDNweDsKCQkJCWRpc3BsYXk6IGlubGluZS1ibG9jazsKCQkJCWN1cnNvcjogcG9pbnRlcjsKCQkJfQoKCQkJLmNyZWF0ZS1mb3JtIHsKCQkJCWJvcmRlcjogc29saWQgZ3JheSAxcHg7CgkJCQlib3JkZXItcmFkaXVzOiA0cHg7CgkJCQliYWNrZ3JvdW5kLWNvbG9yOiAjRjhGOEY4OwoJCQkJcGFkZGluZzogMTZweDsKCQkJfQoKCQkJI2NyZWF0ZS1mb3JtLXRpdGxlIHsKCQkJCWRpc3BsYXk6IGJsb2NrOwoJCQkJd2lkdGg6IDEwMCU7CgkJCQlmb250LXdlaWdodDogYm9sZDsKCQkJCWZvbnQtc2l6ZTogMTUwJTsKCQkJfQoKCQkJI2NyZWF0ZS1mb3JtLWNvbnRlbnQgewoJCQkJZGlzcGxheTogYmxvY2s7CgkJCQl3aWR0aDogMTAwJTsKCQkJfQoKCQkJLmF1dGggewoJCQkJdGV4dC1hbGlnbjogcmlnaHQ7CgkJCX0KCgkJPC9zdHlsZT4KCgkJPHNjcmlwdD4KCQkJZnVuY3Rpb24gcmVtb3ZlQXJ0aWNsZShpZCkgewoJCQkJdmFyIHhociA9IG5ldyBYTUxIdHRwUmVxdWVzdCgpOwoJCQkJeGhyLm9wZW4oJ0RFTEVURScsICcvYXJ0aWNsZXMvJytpZCwgdHJ1ZSk7CgkJCQl4aHIub25sb2FkID0gZnVuY3Rpb24oKSB7CgkJCQkJZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoaWQpLnN0eWxlLmRpc3BsYXkgPSAnbm9uZSc7CgkJCQl9OwoJCQkJeGhyLnNlbmQobnVsbCk7CgkJCX0KCgkJCWZ1bmN0aW9uIGNyZWF0ZUFydGljbGUoKSB7CgoJCQkJdmFyIHhociA9IG5ldyBYTUxIdHRwUmVxdWVzdCgpOwoJCQkJeGhyLm9wZW4oJ1BPU1QnLCAnL2FydGljbGVzJywgdHJ1ZSk7CgkJCQl4aHIub25sb2FkID0gZnVuY3Rpb24oKSB7CgkJCQkJd2luZG93LmxvY2F0aW9uLmhyZWYgPSAnLyc7CgkJCQl9OwoKCQkJCXZhciB0aXRsZSA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKCJjcmVhdGUtZm9ybS10aXRsZSIpOwoJCQkJdmFyIGNvbnRlbnQgPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZCgiY3JlYXRlLWZvcm0tY29udGVudCIpOwoKCQkJCXZhciBwYXlsb2FkID0gewoJCQkJCSJ0aXRsZSI6IHRpdGxlLnZhbHVlLAoJCQkJCSJjb250ZW50IjogY29udGVudC52YWx1ZSwKCQkJCX07CgoJCQkJeGhyLnNlbmQoSlNPTi5zdHJpbmdpZnkocGF5bG9hZCkpOwoJCQl9CgoJCQlmdW5jdGlvbiBsb2dpbihlKSB7CgkJCQllLnByZXZlbnREZWZhdWx0KCk7CgoJCQkJdmFyIGVtYWlsID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoImxvZ2luLWVtYWlsIik7CgkJCQl2YXIgcGFzc3dvcmQgPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZCgibG9naW4tcGFzc3dvcmQiKTsKCgkJCQl2YXIgeGhyID0gbmV3IFhNTEh0dHBSZXF1ZXN0KCk7CgkJCQl4aHIub3BlbignUE9TVCcsICcvbG9naW4vZW1haWwnLCB0cnVlKTsKCQkJCXhoci5vbmxvYWQgPSBmdW5jdGlvbigpIHsKCQkJCQl3aW5kb3cubG9jYXRpb24uaHJlZiA9ICcvJzsKCQkJCX07CgoJCQkJdmFyIHBheWxvYWQgPSB7CgkJCQkJZW1haWw6IGVtYWlsLnZhbHVlLAoJCQkJCXBhc3N3b3JkOiBwYXNzd29yZC52YWx1ZSwKCQkJCX07CgoJCQkJeGhyLnNlbmQoSlNPTi5zdHJpbmdpZnkocGF5bG9hZCkpOwoJCQl9CgoJCQlmdW5jdGlvbiBsb2dvdXQoZSkgewoJCQkJZS5wcmV2ZW50RGVmYXVsdCgpOwoKCQkJCXZhciB4aHIgPSBuZXcgWE1MSHR0cFJlcXVlc3QoKTsKCQkJCXhoci5vcGVuKCdERUxFVEUnLCAnL3Nlc3Npb25zL2N1cnJlbnQnLCB0cnVlKTsKCQkJCXhoci5vbmxvYWQgPSBmdW5jdGlvbigpIHsKCQkJCQl3aW5kb3cubG9jYXRpb24uaHJlZiA9ICcvJzsKCQkJCX07CgoJCQkJeGhyLnNlbmQoKTsKCQkJfQoKCQk8L3NjcmlwdD4KCQl7e2lmIC5nb29nbGVfYW5hbHl0aWNzIH19CgkJPHNjcmlwdCBhc3luYyBzcmM9Imh0dHBzOi8vd3d3Lmdvb2dsZXRhZ21hbmFnZXIuY29tL2d0YWcvanM/aWQ9e3suZ29vZ2xlX2FuYWx5dGljc319Ij48L3NjcmlwdD4KCQk8c2NyaXB0PgoJCSAgd2luZG93LmRhdGFMYXllciA9IHdpbmRvdy5kYXRhTGF5ZXIgfHwgW107CgkJICBmdW5jdGlvbiBndGFnKCl7ZGF0YUxheWVyLnB1c2goYXJndW1lbnRzKTt9CgkJICBndGFnKCdqcycsIG5ldyBEYXRlKCkpOwoJCSAgZ3RhZygnY29uZmlnJywgJ3t7Lmdvb2dsZV9hbmFseXRpY3N9fScpOwoJCTwvc2NyaXB0PgoJCXt7ZW5kfX0KCTwvaGVhZD4KCTxib2R5PgoJCTxkaXYgY2xhc3M9ImNvbnRlbnQiPgoKCQkJPGRpdiBjbGFzcz0iYXV0aCI+CgkJCQl7e2lmIC51c2VyfX0KCQkJCTxmb3JtIGlkPSJmb3JtLWxvZ291dCI+CgkJCQkJe3tpZiAudXNlci5Mb2dpbkdvb2dsZSB9fQoJCQkJCTxpbWcgc3JjPSJ7ey51c2VyLkxvZ2luR29vZ2xlLlBpY3R1cmV9fSIgc3R5bGU9ImhlaWdodDogMjRweDsgdmVydGljYWwtYWxpZ246IHRleHQtYm90dG9tOyI+CgkJCQkJe3tlbmR9fQoJCQkJCXt7LnVzZXIuTmlja319IDxidXR0b24+U2FsaXI8L2J1dHRvbj4KCQkJCTwvZm9ybT4KCQkJCTxzY3JpcHQ+CgkJCQkJZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoJ2Zvcm0tbG9nb3V0JykuYWRkRXZlbnRMaXN0ZW5lcignc3VibWl0JywgbG9nb3V0LCB0cnVlKTsKCQkJCTwvc2NyaXB0PgoJCQkJe3tlbHNlfX0KCQkJCTxmb3JtIGlkPSJmb3JtLWxvZ2luIj4KCQkJCQk8aW5wdXQgdHlwZT0idGV4dCIgaWQ9ImxvZ2luLWVtYWlsIiBwbGFjZWhvbGRlcj0idHVAZW1haWwuY29tIj4KCQkJCQk8aW5wdXQgdHlwZT0icGFzc3dvcmQiIGlkPSJsb2dpbi1wYXNzd29yZCIgcGxhY2Vob2xkZXI9ImNvbnRyYXNlw7FhIj4KCQkJCQk8YnV0dG9uPkVudHJhcjwvYnV0dG9uPgoJCQkJPC9mb3JtPgoJCQkJPGEgaHJlZj0ie3sgLmdvb2dsZV9vYXV0aF9saW5rIH19Ij5FbnRyYXIgY29uIEdvb2dsZTwvYT4KCQkJCTxzY3JpcHQ+CgkJCQkJZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoJ2Zvcm0tbG9naW4nKS5hZGRFdmVudExpc3RlbmVyKCdzdWJtaXQnLCBsb2dpbiwgdHJ1ZSk7CgkJCQk8L3NjcmlwdD4KCQkJCXt7ZW5kfX0KCQkJPC9kaXY+CgoJCQk8YSBocmVmPSIvIj48aDE+QmxvR288L2gxPjwvYT4KCgkJCXt7YmxvY2sgImNvbnRlbnQiIC59fQoJCQkJVEhJUyBJUyBUSEUgQ09OVEVOVAoJCQl7e2VuZH19CgkJPGRpdj4KCTwvYm9keT4KPC9odG1sPgo=`,
}