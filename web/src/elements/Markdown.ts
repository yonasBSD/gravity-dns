import { CSSResult, TemplateResult, html } from "lit";
import { customElement, property } from "lit/decorators.js";
import { unsafeHTML } from "lit/directives/unsafe-html.js";

import PFContent from "@patternfly/patternfly/components/Content/content.css";
import PFList from "@patternfly/patternfly/components/List/list.css";

import { AKElement } from "./Base";

export interface MarkdownDocument {
    html: string;
    metadata: { [key: string]: string };
    filename: string;
}

@customElement("ak-markdown")
export class Markdown extends AKElement {
    @property({ attribute: false })
    md?: MarkdownDocument;

    static get styles(): CSSResult[] {
        return [PFList, PFContent, AKElement.GlobalStyle];
    }

    render(): TemplateResult {
        if (!this.md) {
            return html``;
        }
        const finalHTML = this.md?.html.replace("<ul>", "<ul class='pf-c-list'>");
        return html`${this.md?.metadata.title ? html`<h2>${this.md.metadata.title}</h2>` : html``}
        ${unsafeHTML(finalHTML)}`;
    }
}
