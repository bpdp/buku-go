all:
	pdflatex pemrograman-go.tex
	makeindex pemrograman-go.idx
	pdflatex pemrograman-go.tex

clean-all:
	rm -rf *.aux; rm *.pdf; rm -rf *.lof; rm -rf *.log; rm -rf *.toc; rm -rf *.lol; rm -rf *.idx; rm -rf *.out; rm -rf *.lot; rm -rf *.tex~; rm -rf *.ind; rm -rf *.ilg; rm -rf components/*.tex~

clean-without-pdf:
	rm -rf *.aux; rm -rf *.lof; rm -rf *.log; rm -rf *.toc; rm -rf *.lol; rm -rf *.idx; rm -rf *.out; rm -rf *.lot; rm -rf *.tex~; rm -rf *.ind; rm -rf *.ilg; rm -rf components/*.tex~

view:
	xpdf pemrograman-go.pdf
