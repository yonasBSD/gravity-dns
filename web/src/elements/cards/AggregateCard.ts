import { CSSResult, TemplateResult, css, html } from "lit";
import { customElement, property } from "lit/decorators.js";
import { ifDefined } from "lit/directives/if-defined.js";

import PFCard from "@patternfly/patternfly/components/Card/card.css";
import PFFlex from "@patternfly/patternfly/layouts/Flex/flex.css";
import PFBase from "@patternfly/patternfly/patternfly-base.css";

import { AKElement } from "../Base";

@customElement("ak-aggregate-card")
export class AggregateCard extends AKElement {
    @property()
    icon?: string;

    @property()
    header?: string;

    @property()
    headerLink?: string;

    @property({ type: Boolean })
    isCenter = true;

    static get styles(): CSSResult[] {
        return [PFBase, PFCard, PFFlex, AKElement.GlobalStyle].concat([
            css`
                .pf-c-card.pf-c-card-aggregate {
                    height: 100%;
                }
                .pf-c-card__header {
                    flex-wrap: nowrap;
                }
                .center-value {
                    font-size: var(--pf-global--icon--FontSize--lg);
                    text-align: center;
                    color: var(--pf-global--Color--100);
                }
                .subtext {
                    font-size: var(--pf-global--FontSize--sm);
                }
            `,
        ]);
    }

    renderInner(): TemplateResult {
        return html`<slot></slot>`;
    }

    renderHeaderLink(): TemplateResult {
        return html`${this.headerLink
            ? html`<a href="${this.headerLink}">
                  <i class="fa fa-link"> </i>
              </a>`
            : ""}`;
    }

    render(): TemplateResult {
        return html`<div class="pf-c-card pf-c-card-aggregate">
            <div class="pf-c-card__header pf-l-flex pf-m-justify-content-space-between">
                <div class="pf-c-card__title">
                    <i class="${ifDefined(this.icon)}"></i>&nbsp;${this.header ? this.header : ""}
                </div>
                ${this.renderHeaderLink()}
            </div>
            <div class="pf-c-card__body ${this.isCenter ? "center-value" : ""}">
                ${this.renderInner()}
            </div>
        </div>`;
    }
}
